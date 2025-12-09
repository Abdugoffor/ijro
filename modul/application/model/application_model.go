package application_model

import (
	category_model "ijro-nazorat/modul/category/model"
	country_model "ijro-nazorat/modul/country/model"
	user_model "ijro-nazorat/modul/user/model"
)

type Application struct {
	ID          int                     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      int                     `json:"user_id"`
	User        user_model.User         `json:"user" gorm:"foreignKey:UserID"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Image       string                  `json:"image"`
	File        string                  `json:"file"`
	CategoryID  int                     `json:"category_id"`
	Category    category_model.Category `json:"category" gorm:"foreignKey:CategoryID"`
	CountryID   int                     `json:"country_id"`
	Country     country_model.Country   `json:"country" gorm:"foreignKey:CountryID"`
	Status      string                  `json:"status" gorm:"default:pending"`
	CreatedAt   string                  `json:"created_at"`
	UpdatedAt   string                  `json:"updated_at"`
}

func (Application) TableName() string {
	return "applications"
}
