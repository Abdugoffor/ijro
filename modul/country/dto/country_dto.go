package country_dto

import (
	"ijro-nazorat/helper"
	country_model "ijro-nazorat/modul/country/model"
)

type CreateOrUpdate struct {
	Name     string `json:"name"`
	IsActive *bool  `json:"is_active"`
}

type Filter struct {
	Name   string `json:"name" query:"name"`
	Status string `json:"status" query:"status"`
	Sort   string `json:"sort" query:"sort"`
	Column string `json:"column" query:"column"`
}

type Response struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func ToResponse(model country_model.Country) Response {
	return Response{
		ID:        model.ID,
		Name:      model.Name,
		IsActive:  model.IsActive,
		CreatedAt: helper.FormatDate(model.CreatedAt),
		UpdatedAt: helper.FormatDate(model.UpdatedAt),
		DeletedAt: helper.FormatDate(model.DeletedAt),
	}
}
