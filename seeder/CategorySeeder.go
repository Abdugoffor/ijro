package seeder

import (
	"ijro-nazorat/config"
	category_model "ijro-nazorat/modul/category/model"
	"log"
)

func CategorySeeder() {
	categories := []category_model.Category{
		{Name: "Electronics", IsActive: true},
		{Name: "Clothing", IsActive: true},
		{Name: "Home", IsActive: true},
		{Name: "Beauty", IsActive: true},
		{Name: "Sports", IsActive: true},
		{Name: "Toys", IsActive: true},
	}

	for _, cat := range categories {
		config.DB.Create(&cat)
	}

	log.Println("✅ CategorySeeder completed")
}
