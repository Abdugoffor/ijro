package config

import (
	category_model "ijro-nazorat/modul/category/model"
	country_model "ijro-nazorat/modul/country/model"
	"log"
)

func RunMigrations() {
	models := []interface{}{
		&category_model.Category{},
		&country_model.Country{},
	}

	err := DB.AutoMigrate(models...)
	{
		if err != nil {
			log.Println("❌ failed to migrate models:", err)
		}
	}

	log.Println("✅ Migrations completed")
}
