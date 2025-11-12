package config

import (
	"fmt"
	"ijro-nazorat/helper"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreSQL
func PostgreSQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		helper.ENV("DB_HOST"),
		helper.ENV("DB_USER"),
		helper.ENV("DB_PASSWORD"),
		helper.ENV("DB_NAME"),
		helper.ENV("DB_PORT"),
		helper.ENV("DB_SSLMODE"),
		helper.ENV("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ failed to connect to PostgreSQL: %w", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	return db, nil
}

// MySQL
func MySQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		helper.ENV("DB_USER"),
		helper.ENV("DB_PASSWORD"),
		helper.ENV("DB_HOST"),
		helper.ENV("DB_PORT"),
		helper.ENV("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ failed to connect to MySQL: %w", err)
	}

	log.Println("✅ Connected to MySQL")
	return db, nil
}

// SQLite
func SQLite() (*gorm.DB, error) {
	sqlitePath := helper.ENV("DB_PATH")
	if sqlitePath == "" {
		sqlitePath = "data.db"
	}

	// Fayl mavjudligini tekshirish
	if _, err := os.Stat(sqlitePath); os.IsNotExist(err) {
		file, createErr := os.Create(sqlitePath)
		if createErr != nil {
			return nil, fmt.Errorf("❌ failed to create SQLite file: %w", createErr)
		}
		file.Close()
		log.Println("🆕 SQLite file created:", sqlitePath)
	}

	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ failed to connect to SQLite: %w", err)
	}

	log.Println("✅ Connected to SQLite (" + sqlitePath + ")")
	return db, nil
}
