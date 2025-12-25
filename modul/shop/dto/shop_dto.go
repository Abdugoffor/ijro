package shop_dto

import (
	"time"

	"git.sriss.uz/shared/shared_service/response"
	"gorm.io/gorm"
)

type ShopPage = response.PageData[ShopResponse]

type ShopChange struct {
	Name     string `json:"name" validate:"required"`
	Slug     string `json:"slug"`
	IsActive *bool  `json:"is_active" validate:"required"`
}

type ShopResponse struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	IsActive  bool           `json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
