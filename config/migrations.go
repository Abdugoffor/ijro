package config

import (
	category_model "ijro-nazorat/modul/category/model"
	country_model "ijro-nazorat/modul/country/model"
	user_model "ijro-nazorat/modul/user/model"
	"log"
)

func RunMigrations() {
	models := []interface{}{
		&category_model.Category{},
		&country_model.Country{},
		&user_model.User{},
	}

	err := DB.AutoMigrate(models...)
	{
		if err != nil {
			log.Println("❌ failed to migrate models:", err)
		}
	}

	log.Println("✅ Migrations completed")
}
