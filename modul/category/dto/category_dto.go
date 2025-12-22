package category_dto

import (
	"ijro-nazorat/helper"
	category_model "ijro-nazorat/modul/category/model"

	"git.sriss.uz/shared/shared_service/response"
)

type CategoryPage = response.PageData[Response]

type CreateOrUpdate struct {
	Name     string `json:"name"`
	IsActive *bool  `json:"is_active"` // pointer bool
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

func ToResponse(category category_model.Category) Response {
	return Response{
		ID:        category.ID,
		Name:      category.Name,
		IsActive:  category.IsActive,
		CreatedAt: helper.FormatDate(category.CreatedAt),
		UpdatedAt: helper.FormatDate(category.UpdatedAt),
		DeletedAt: helper.FormatDate(category.DeletedAt),
	}
}
