package user_model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primary_key"`
	Name      string         `json:"name" gorm:"name"`
	Email     string         `json:"email" gorm:"uniqueIndex"`
	Password  string         `json:"password"`
	Role      string         `json:"role" default:"user"`
	CountryID *int           `json:"country_id" gorm:"column:country_id"`
	IsActive  bool           `json:"is_active" default:"true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
