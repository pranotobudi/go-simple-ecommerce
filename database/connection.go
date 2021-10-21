package database

import (
	"fmt"
	"log"

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
	fmt.Println("database connection is success..")
	return db
}
