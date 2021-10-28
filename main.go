package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-simple-ecommerce/api/products"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("failed to load .env file")
	}

	// db := database.GetDBInstance()
	pr := products.NewProductRepository()
	pr.FreshProductMigrator()
	pr.ProductDataSeed()
	productsList, _ := pr.GetProducts()
	fmt.Println(productsList)

	handler := products.NewProductHandler()
	e := echo.New()
	e.GET("/", hello)
	e.GET("/api/v1/products", handler.GetProducts)
	e.Static("/static", "assets")

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Go World!")
}
