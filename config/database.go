package config

import (
	"log"

	"ijro-nazorat/helper"

	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() *gorm.DB {
	driver := helper.ENV("DB_DRIVER")

	var db *gorm.DB

	var err error

	switch driver {

	case "postgres":
		db, err = PostgreSQL()

	case "mysql":
		db, err = MySQL()

	case "sqlite":
		db, err = SQLite()

	default:
		log.Println("⚠️  Unknown DB_DRIVER, defaulting to SQLite")
		db, err = SQLite()
	}

	if err != nil {
		log.Printf("❌ failed to connect to %s DB: %v", driver, err)
	}

	DB = db

	RunMigrations()

	return db
}
