package products

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `json:"title"`
	Price       string `json:"price"`
	Rating      string `json:"rating"`
	Description string `json:"description"`
	Category    string `json:"category"`
	ImageUrl    string `json:"image_url"`
}
