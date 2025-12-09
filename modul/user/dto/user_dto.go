package user_dto

import (
	"ijro-nazorat/helper"
	user_model "ijro-nazorat/modul/user/model"
)

type CreateOrUpdate struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CountryID int    `json:"country_id"`
	IsActive  *bool  `json:"is_active"`
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

type Response struct {
	ID        int    `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`
	Email     string `json:"email" gorm:"column:email"`
	Role      string `json:"role" gorm:"column:role"`
	Country   string `json:"country" gorm:"column:country"` // ← country
	IsActive  bool   `json:"is_active" gorm:"column:is_active"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
}

func ToResponse(user user_model.User) Response {
	return Response{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Country:   user.Country.Name,
		IsActive:  user.IsActive,
		CreatedAt: helper.FormatDate(user.CreatedAt),
	}
}
