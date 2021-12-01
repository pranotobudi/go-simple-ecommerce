package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-simple-ecommerce/api/products"
	"github.com/pranotobudi/go-simple-ecommerce/api/users"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		// executed in development only
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("failed to load .env file")
		}
	}

	// Product
	// db := database.GetDBInstance()
	pr := products.NewProductRepository()
	pr.FreshProductMigrator()
	pr.ProductDataSeed()
	productsList, _ := pr.GetProducts()
	fmt.Println(productsList)

	//User
	ur := users.NewUserRepository()
	ur.FreshUserMigrator()
	ur.UserDataSeed()
	user, _ := ur.GetUserByEmail("emailusername1@gmail.com")
	fmt.Println(user)

	productHandler := products.NewProductHandler()
	userHandler := users.NewUserHandler()
	e := echo.New()
	e.GET("/", hello)
	e.GET("/api/v1/products", productHandler.GetProducts)
	e.POST("/api/v1/register", userHandler.RegisterUser)
	e.POST("/api/v1/login", userHandler.UserLogin)
	e.Static("/static", "assets")

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Go World!")
}
