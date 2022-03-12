package products

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-simple-ecommerce/api"
)

type SearchRequest struct {
	SearchTerm string `json:"search_term"`
}

type ProductResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
	Category    string `json:"category"`
	ImageUrl    string `json:"image"`
}

func ProductsResponseFormatter(products []Product) []ProductResponse {
	var formatters []ProductResponse
	for _, product := range products {
		formatter := ProductResponse{
			ID:          product.ID,
			Title:       product.Title,
			Price:       product.Price,
			Rating:      product.Rating,
			Description: product.Description,
			Category:    product.Category,
			ImageUrl:    product.ImageUrl,
		}
		formatters = append(formatters, formatter)
	}
	return formatters
}

type productHandler struct {
	repository ProductRepository
}

func NewProductHandler() *productHandler {
	repository := NewProductRepository()

	return &productHandler{repository}
}

func (h *productHandler) GetProducts(c echo.Context) error {
	// Get Products from repository
	products, err := h.repository.GetProducts()
	if err != nil {
		return api.ResponseErrorFormatter(c, err)
	}

	// Success ProductResponse
	data := ProductsResponseFormatter(products)

	response := api.ResponseFormatter(http.StatusOK, "success", "get products successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *productHandler) SearchProducts(c echo.Context) error {
	// Input Binding
	searchReq := new(SearchRequest)
	if err := c.Bind(searchReq); err != nil {
		return api.ResponseErrorFormatter(c, err)
	}
	fmt.Println("=========Search Request: ", searchReq)

	// Get Products from Elastic Search repository
	products, err := h.repository.SearchProducts(searchReq.SearchTerm)
	if err != nil {
		return api.ResponseErrorFormatter(c, err)
	}
	// products := []Product{{Title: "product1"}, {Title: "product2"}}

	// Success ProductResponse
	data := ProductsResponseFormatter(products)

	response := api.ResponseFormatter(http.StatusOK, "success", "get products successfull", data)

	return c.JSON(http.StatusOK, response)
}
