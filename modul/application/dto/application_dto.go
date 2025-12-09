package application_dto

import application_model "ijro-nazorat/modul/application/model"

type Create struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	File        string `json:"file"`
	CategoryID  int    `json:"category_id"`
	CountryID   int    `json:"country_id"` // to'g'ri nom
}

type Update struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	File        string `json:"file"`
	CategoryID  int    `json:"category_id"`
	CountryID   int    `json:"country_id"` // to'g'ri nom
}

type Response struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	User        string `json:"user"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	File        string `json:"file"`
	CategoryID  int    `json:"category_id"`
	Category    string `json:"category"`
	CountryID   int    `json:"country_id"` // to'g'ri nom
	Country     string `json:"country"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Filter struct {
	Name      string `json:"name" query:"name"`
	Email     string `json:"email" query:"email"`
	Role      string `json:"role" query:"role"`
	CountryID *int   `json:"country_id" query:"country_id"`
	Status    string `json:"status" query:"status"`
	Sort      string `json:"sort" query:"sort"`
	Column    string `json:"column" query:"column"`
}

func ToResponse(application application_model.Application) Response {
	return Response{
		ID:          application.ID,
		UserID:      application.UserID,
		User:        application.User.Name,
		Name:        application.Name,
		Description: application.Description,
		Image:       application.Image,
		File:        application.File,
		CategoryID:  application.CategoryID,
		Category:    application.Category.Name,
		CountryID:   application.CountryID,
		Country:     application.Country.Name,
		Status:      application.Status,
		CreatedAt:   application.CreatedAt,
		UpdatedAt:   application.UpdatedAt,
	}
}
