package helper

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ .env fayl topilmadi yoki yuklanmadi")
	}
}

func ENV(key string) string {
	return os.Getenv(key)
}

func FormatDate(v any) string {
	switch t := v.(type) {
	case time.Time:
		return t.Format("02-01-2006 15:04:05")
	case gorm.DeletedAt:
		if t.Valid {
			return t.Time.Format("02-01-2006 15:04:05")
		}
		return ""
	default:
		return ""
	}
}
