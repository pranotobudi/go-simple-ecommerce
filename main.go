package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	// "github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-simple-ecommerce/internal/repository"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("failed to load .env file")
	}

	// db := database.GetDBInstance()
	pr := repository.NewProductRepository()

	pr.FreshProductMigrator()
	pr.ProductDataSeed()
	products, _ := pr.GetProducts()
	fmt.Println(products)
	// e := echo.New()
	// e.GET("/", hello)
	// e.Logger.Fatal(e.Start(":8080"))
}

// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, Go World!")
// }
