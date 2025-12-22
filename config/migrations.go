package config

import (
	application_model "ijro-nazorat/modul/application/model"
	category_model "ijro-nazorat/modul/category/model"
	country_model "ijro-nazorat/modul/country/model"
	form_model "ijro-nazorat/modul/form/model"
	user_model "ijro-nazorat/modul/user/model"
	"log"
)

func RunMigrations() {
	models := []interface{}{
		&category_model.Category{},
		&country_model.Country{},
		&user_model.User{},
		&application_model.Application{},
		&application_model.Answer{},
		&form_model.AppCategory{},
		&form_model.Page{},
		&form_model.Form{},
		&form_model.App{},
		&form_model.AppInfo{},
	}

	err := DB.AutoMigrate(models...)
	{
		if err != nil {
			log.Println("❌ failed to migrate models:", err)
		}
	}

	log.Println("✅ Migrations completed")
}
