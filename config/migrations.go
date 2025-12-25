package config

import (
	category_model "ijro-nazorat/modul/category/model"
	shop_model "ijro-nazorat/modul/shop/model"
	"log"
)

func RunMigrations() {
	models := []interface{}{
		&category_model.Category{},
		&shop_model.Shop{},
		&shop_model.ShopItems{},
		&shop_model.ShopItemsLog{},
	}

	err := DB.AutoMigrate(models...)
	{
		if err != nil {
			log.Println("❌ failed to migrate models:", err)
		}
	}

	log.Println("✅ Migrations completed")
}
