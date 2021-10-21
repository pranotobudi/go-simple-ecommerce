package main

import (
	"log"

	"github.com/joho/godotenv"
	// "github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-simple-ecommerce/database"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("failed to load .env file")
	}

	database.GetDBInstance()
	// e := echo.New()
	// e.GET("/", hello)
	// e.Logger.Fatal(e.Start(":8080"))
}

// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, Go World!")
// }
