package repository

import (
	"fmt"
	"time"

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
	ProductSeeder()
	AddProduct(product Product) (*Product, error)
	GetProduct(id uint) (*Product, error)
	GetProducts() ([]Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository() *productRepository {
	db := database.GetDBInstance()
	return &productRepository{db}
}

func (r *productRepository) FreshProductMigrator() {
	r.db.AutoMigrate(Product{})

	// Create Fresh Recipe Table
	if (r.db.Migrator().HasTable(&Product{})) {
		fmt.Println("Product table exist")
		r.db.Migrator().DropTable(&Product{})
		fmt.Println("Drop Product table")
	}
	r.db.Migrator().CreateTable(&Product{})
	fmt.Println("Create Product table")

}

func (r *productRepository) ProductDataSeed() {
	statement := "INSERT INTO products (title, price, rating, description, category, image_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	r.db.Exec(statement, "title1", 7, 5, "description 1", "category 1", "/assets/images/product01.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title2", 7, 4, "description 2", "category 2", "/assets/images/product02.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title3", 7, 3, "description 3", "category 3", "/assets/images/product03.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title4", 7, 3, "description 4", "category 1", "/assets/images/product04.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title5", 7, 4, "description 5", "category 2", "/assets/images/product05.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title6", 7, 5, "description 6", "category 3", "/assets/images/product06.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title7", 7, 2, "description 7", "category 1", "/assets/images/product07.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title8", 7, 1, "description 8", "category 2", "/assets/images/product08.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title9", 7, 5, "description 9", "category 3", "/assets/images/product09.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title10", 7, 3, "description 10", "category 1", "/assets/images/product10.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title11", 7, 5, "description 11", "category 2", "/assets/images/product11.jpg", time.Now(), time.Now())
	r.db.Exec(statement, "title12", 7, 2, "description 12", "category 3", "/assets/images/product12.jpg", time.Now(), time.Now())
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
