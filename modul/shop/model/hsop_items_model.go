package shop_model

import (
	"time"

	"gorm.io/gorm"
)

type ShopItems struct {
	ID        int64          `json:"id"`
	ShopID    int64          `json:"shop_id"`
	ProductID int64          `json:"product_id"`
	Value     string         `json:"value"`
	Unit      string         `json:"unit"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ShopItems) TableName() string {
	return "shop_items"
}
