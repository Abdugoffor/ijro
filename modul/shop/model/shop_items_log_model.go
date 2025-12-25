package shop_model

import (
	"time"

	"gorm.io/gorm"
)

type ShopItemsLog struct {
	ID        int64          `json:"id"`
	ShopID    int64          `json:"shop_id"`
	ProductID int64          `json:"product_id"`
	Value     string         `json:"value"`
	Unit      string         `json:"unit"`
	Price     int64          `json:"price"`
	OldPrice  int64          `json:"old_price"`
	Type      string         `json:"type"`       // prixod, rasxod
	FromWhere string         `json:"from_where"` // ombordan, shopdan, sotuvdan, prixoddan
	WhereId   int64          `json:"where_id"`   // Qayerga ketdi , qayerdan keldi, shop, wherehous
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ShopItemsLog) TableName() string {
	return "shop_items_log"
}
