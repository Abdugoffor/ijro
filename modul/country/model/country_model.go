package country_model

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	IsActive  bool           `json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Country) TableName() string {
	return "countries"
}
