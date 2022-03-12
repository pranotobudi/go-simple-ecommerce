package products

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pranotobudi/go-simple-ecommerce/config"
	"github.com/pranotobudi/go-simple-ecommerce/database"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	Price       int
	Rating      int
	Description string
	Category    string
	ImageUrl    string
}

type ProductRepository interface {
	FreshProductMigrator()
	FreshProductMigratorPostgres()
	FreshProductMigratorElasticSearch()
	ProductDataSeed()
	ProductDataSeedPostgres()
	AddProduct(product Product) (*Product, error)
	GetProduct(id uint) (*Product, error)
	GetProducts() ([]Product, error)
	SearchProducts(searchTerm string) ([]Product, error)
}

type productRepository struct {
	db *gorm.DB
	es *elasticsearch.Client
}

func NewProductRepository() *productRepository {
	db := database.GetDBInstance()
	es := database.GetElasticSearchDBInstance()
	return &productRepository{
		db: db,
		es: es,
	}
}

func (r *productRepository) FreshProductMigrator() {
	r.FreshProductMigratorPostgres()
	r.FreshProductMigratorElasticSearch()
}

func (r *productRepository) FreshProductMigratorPostgres() {
	r.db.AutoMigrate(Product{})

	// Create Fresh Product Table
	if (r.db.Migrator().HasTable(&Product{})) {
		fmt.Println("Postgres Product table exist")
		r.db.Migrator().DropTable(&Product{})
		fmt.Println("Drop Postgres Product table")
	}
	r.db.Migrator().CreateTable(&Product{})
	fmt.Println("Create Postgres Product table")
}
func (r *productRepository) FreshProductMigratorElasticSearch() {
	// TODO: implement this
	esDbConfig := config.ESDbConfig()

	query := `{
		"query": {
		  "match_all": {}
		}
	  }`

	reader := strings.NewReader(query)

	r.es.DeleteByQuery([]string{esDbConfig.IndexName}, reader)
	fmt.Printf("Drop ElasticSearch %s index documents", esDbConfig.IndexName)
}

func (r *productRepository) ProductDataSeed() {
	r.ProductDataSeedPostgres()
	r.ProductDataSeedElasticSearch()
}
func (r *productRepository) ProductDataSeedPostgres() {
	statement := "INSERT INTO products (title, price, rating, description, category, image_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	appConfig := config.AppConfig()

	r.db.Exec(statement, "title1", 7, 5, "backend description 1", "category 1", appConfig.AppHost+"/static/images/product01.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title2", 8, 4, "backend description 2", "category 2", appConfig.AppHost+"/static/images/product02.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title3", 6, 3, "backend description 3", "category 3", appConfig.AppHost+"/static/images/product03.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title4", 5, 3, "backend description 4", "category 1", appConfig.AppHost+"/static/images/product04.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title5", 9, 4, "backend description 5", "category 2", appConfig.AppHost+"/static/images/product05.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title6", 7, 5, "backend description 6", "category 3", appConfig.AppHost+"/static/images/product06.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title7", 9, 2, "backend description 7", "category 1", appConfig.AppHost+"/static/images/product07.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title8", 8, 1, "backend description 8", "category 2", appConfig.AppHost+"/static/images/product08.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title9", 5, 5, "backend description 9", "category 3", appConfig.AppHost+"/static/images/product09.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title10", 7, 3, "backend description 10", "category 1", appConfig.AppHost+"/static/images/product10.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title11", 9, 5, "backend description 11", "category 2", appConfig.AppHost+"/static/images/product11.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title12", 7, 2, "backend description 12", "category 3", appConfig.AppHost+"/static/images/product12.jpg", time.Now(), time.Now())
}

func (r *productRepository) ProductDataSeedElasticSearch() {
	appConfig := config.AppConfig()
	esDbConfig := config.ESDbConfig()

	queries := []string{
		`{
			"title": "title1",
			"price":"7",
			"rating": "5",
			"description": "backend description 1",
			"category":"category 1",
			"image_url": "` + appConfig.AppHost + `/static/images/product01.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title2",
			"price":"8",
			"rating": "4",
			"description": "backend description 2",
			"category":"category 2",
			"image_url": "` + appConfig.AppHost + `/static/images/product02.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title3",
			"price":"6",
			"rating": "3",
			"description": "backend description 3",
			"category":"category 3",
			"image_url": "` + appConfig.AppHost + `/static/images/product03.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title4",
			"price":"5",
			"rating": "3",
			"description": "backend description 4",
			"category":"category 1",
			"image_url": "` + appConfig.AppHost + `/static/images/product04.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title5",
			"price":"9",
			"rating": "4",
			"description": "backend description 5",
			"category":"category 2",
			"image_url": "` + appConfig.AppHost + `/static/images/product05.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title6",
			"price":"7",
			"rating": "5",
			"description": "backend description 6",
			"category":"category 3",
			"image_url": "` + appConfig.AppHost + `/static/images/product06.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title7",
			"price":"9",
			"rating": "2",
			"description": "backend description 7",
			"category":"category 1",
			"image_url": "` + appConfig.AppHost + `/static/images/product07.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title8",
			"price":"8",
			"rating": "1",
			"description": "backend description 8",
			"category":"category 2",
			"image_url": "` + appConfig.AppHost + `/static/images/product08.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title9",
			"price":"5",
			"rating": "5",
			"description": "backend description 9",
			"category":"category 3",
			"image_url": "` + appConfig.AppHost + `/static/images/product09.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title10",
			"price":"7",
			"rating": "3",
			"description": "backend description 10",
			"category":"category 1",
			"image_url": "` + appConfig.AppHost + `/static/images/product10.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title11",
			"price":"9",
			"rating": "5",
			"description": "backend description 11",
			"category":"category 2",
			"image_url": "` + appConfig.AppHost + `/static/images/product11.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
		`{
			"title": "title12",
			"price":"7",
			"rating": "2",
			"description": "backend description 12",
			"category":"category 3",
			"image_url": "` + appConfig.AppHost + `/static/images/product12.jpg",
			"created_at": "` + time.Now().String() + `",
			"updated_at": "` + time.Now().String() + `"
		}`,
	}
	for idx, query := range queries {
		reader := strings.NewReader(query)
		r.es.Create(esDbConfig.IndexName, strconv.Itoa(idx+1), reader)
	}
	fmt.Printf("Create ElasticSearch %s index data seeds", esDbConfig.IndexName)

	// Perform the search request.
	// ElasticSearch doc index we can define manually
}

func (r *productRepository) AddProduct(entity Product) (*Product, error) {
	err := r.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	err = r.db.First(&entity).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("INSIDE REPOSITORY AddEntity: %+v \n", entity)
	return &entity, nil
}

func (r *productRepository) GetProduct(id uint) (*Product, error) {
	var entity Product
	err := r.db.First(&entity, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *productRepository) GetProducts() ([]Product, error) {
	var entities []Product
	err := r.db.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *productRepository) SearchProducts(searchTerm string) ([]Product, error) {
	var entities []Product
	esDbConfig := config.ESDbConfig()

	// 3. Search for the indexed documents
	//
	// Build the request body.
	var buf bytes.Buffer

	// get by multi field and partial search
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": "*" + searchTerm + "*",
			},
		},
	}

	// // get by multi fields, exact match
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"query_string": map[string]interface{}{
	// 			"query": searchTerm,
	// 		},
	// 	},
	// }

	// // get by field
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"match": map[string]interface{}{
	// 			"title": searchTerm,
	// 		},
	// 	},
	// }

	// // get all
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"match_all": map[string]interface{}{},
	// 	},
	// }

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.

	res, err := r.es.Search(
		r.es.Search.WithContext(context.Background()),
		r.es.Search.WithIndex(esDbConfig.IndexName),
		r.es.Search.WithBody(&buf),
		r.es.Search.WithTrackTotalHits(true),
		r.es.Search.WithPretty(),
		r.es.Search.WithSize(20), // default is 10
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var body map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(body["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(body["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range body["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		id, _ := strconv.Atoi(fmt.Sprintf("%s", hit.(map[string]interface{})["_id"]))
		source := hit.(map[string]interface{})["_source"]
		title := source.(map[string]interface{})["title"]
		price := source.(map[string]interface{})["price"]
		rating := source.(map[string]interface{})["rating"]
		description := source.(map[string]interface{})["description"]
		category := source.(map[string]interface{})["category"]
		image_url := source.(map[string]interface{})["image_url"]
		// log.Println("== TITLE: ", title, price, rating, description, category, image_url)

		entity := Product{}
		entity.ID = uint(id)
		entity.Title = title.(string)
		intPrice, _ := strconv.Atoi(price.(string))
		entity.Price = intPrice
		intRating, _ := strconv.Atoi(rating.(string))
		entity.Rating = intRating
		entity.Description = description.(string)
		entity.Category = category.(string)
		entity.ImageUrl = image_url.(string)

		entities = append(entities, entity)
	}

	log.Println(strings.Repeat("=", 37))
	return entities, nil
}

// func (r *productRepository) SearchProducts(searchTerm string) ([]Product, error) {
// 	var entities []Product

// 	// 3. Search for the indexed documents
// 	//
// 	// Build the request body.
// 	var buf bytes.Buffer
// 	query := map[string]interface{}{
// 		"query": map[string]interface{}{
// 			"match": map[string]interface{}{
// 				"first_name": "budi",
// 			},
// 		},
// 	}
// 	if err := json.NewEncoder(&buf).Encode(query); err != nil {
// 		log.Fatalf("Error encoding query: %s", err)
// 	}

// 	// Perform the search request.
// 	res, err := r.es.Search(
// 		r.es.Search.WithContext(context.Background()),
// 		r.es.Search.WithIndex("learning_es"),
// 		r.es.Search.WithBody(&buf),
// 		r.es.Search.WithTrackTotalHits(true),
// 		r.es.Search.WithPretty(),
// 	)
// 	if err != nil {
// 		log.Fatalf("Error getting response: %s", err)
// 	}
// 	defer res.Body.Close()

// 	if res.IsError() {
// 		var e map[string]interface{}
// 		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
// 			log.Fatalf("Error parsing the response body: %s", err)
// 		} else {
// 			// Print the response status and error information.
// 			log.Fatalf("[%s] %s: %s",
// 				res.Status(),
// 				e["error"].(map[string]interface{})["type"],
// 				e["error"].(map[string]interface{})["reason"],
// 			)
// 		}
// 	}

// 	var body map[string]interface{}

// 	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
// 		log.Fatalf("Error parsing the response body: %s", err)
// 	}
// 	// Print the response status, number of results, and request duration.
// 	log.Printf(
// 		"[%s] %d hits; took: %dms",
// 		res.Status(),
// 		int(body["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
// 		int(body["took"].(float64)),
// 	)
// 	// Print the ID and document source for each hit.
// 	for _, hit := range body["hits"].(map[string]interface{})["hits"].([]interface{}) {
// 		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
// 	}

// 	log.Println(strings.Repeat("=", 37))
// 	return entities, nil
// }

// func (r *productRepository) SearchProducts(searchTerm string) ([]Product, error) {
// 	var entities []Product

// 	esDbConfig := config.ESDbConfig()
// 	query := `{
// 		"query": {
// 			"match_all": {}
// 		}
// 	}`

// 	reader := strings.NewReader(query)

// 	// Perform the search request.

// 	res, err := r.es.Search(
// 		r.es.Search.WithContext(context.Background()),
// 		r.es.Search.WithIndex(esDbConfig.IndexName),
// 		r.es.Search.WithBody(reader),
// 		r.es.Search.WithTrackTotalHits(true),
// 		r.es.Search.WithPretty(),
// 		r.es.Search.WithSize(20), // default is 10
// 	)
// 	if err != nil {
// 		log.Fatalf("Error getting response: %s", err)
// 	}
// 	defer res.Body.Close()

// 	if res.IsError() {
// 		var e map[string]interface{}
// 		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
// 			log.Fatalf("Error parsing the response body: %s", err)
// 		} else {
// 			// Print the response status and error information.
// 			log.Fatalf("[%s] %s: %s",
// 				res.Status(),
// 				e["error"].(map[string]interface{})["type"],
// 				e["error"].(map[string]interface{})["reason"],
// 			)
// 		}
// 	}

// 	var body map[string]interface{}

// 	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
// 		log.Fatalf("Error parsing the response body: %s", err)
// 	}
// 	// Print the response status, number of results, and request duration.
// 	log.Printf(
// 		"[%s] %d hits; took: %dms",
// 		res.Status(),
// 		int(body["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
// 		int(body["took"].(float64)),
// 	)
// 	// Print the ID and document source for each hit.
// 	for _, hit := range body["hits"].(map[string]interface{})["hits"].([]interface{}) {
// 		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
// 		id, _ := strconv.Atoi(fmt.Sprintf("%s", hit.(map[string]interface{})["_id"]))
// 		source := hit.(map[string]interface{})["_source"]
// 		title := source.(map[string]interface{})["title"]
// 		price := source.(map[string]interface{})["price"]
// 		rating := source.(map[string]interface{})["rating"]
// 		description := source.(map[string]interface{})["description"]
// 		category := source.(map[string]interface{})["category"]
// 		image_url := source.(map[string]interface{})["image_url"]
// 		// log.Println("== TITLE: ", title, price, rating, description, category, image_url)

// 		entity := Product{}
// 		entity.ID = uint(id)
// 		entity.Title = title.(string)
// 		intPrice, _ := strconv.Atoi(price.(string))
// 		entity.Price = intPrice
// 		intRating, _ := strconv.Atoi(rating.(string))
// 		entity.Rating = intRating
// 		entity.Description = description.(string)
// 		entity.Category = category.(string)
// 		entity.ImageUrl = image_url.(string)

// 		entities = append(entities, entity)
// 	}

// 	log.Println(strings.Repeat("=", 37))
// 	return entities, nil
// }
