package category_dto

import (
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
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
