package category_model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	IsActive  bool           `json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Category) TableName() string {
	return "categories"
}
