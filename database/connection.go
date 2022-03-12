package database

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pranotobudi/go-simple-ecommerce/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBInstance() *gorm.DB {
	dbConfig := config.DbConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName, dbConfig.Port, dbConfig.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to make database connection...")
	}
	fmt.Println("postgres connection is success..")
	return db
}

func GetElasticSearchDBInstance() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://3ijwfwuxpf:p7xye9ej38@elasticsearch-testin-7721560410.us-east-1.bonsaisearch.net:443",
			// "http://localhost:9201",
		},
		// ServiceToken: "3ijwfwuxpf:p7xye9ej38",
		// APIKey: "3ijwfwuxpf:p7xye9ej38",
		// Username: "3ijwfwuxpf",
		// Password: "p7xye9ej38",
		// ServiceToken: "https://3ijwfwuxpf:p7xye9ej38@elasticsearch-testin-7721560410.us-east-1.bonsaisearch.net:443",
		// ...

	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("elastic search connection failed..: ", err)
		return nil
	}
	log.Println("Elastic Search Version: ", elasticsearch.Version)
	log.Println(es.Info())
	fmt.Println("elastic search connection success..")
	return es
}
