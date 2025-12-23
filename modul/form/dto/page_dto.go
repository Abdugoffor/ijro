package form_dto

import (
	"time"

	"git.sriss.uz/shared/shared_service/response"
	"git.sriss.uz/shared/shared_service/sharedutil"
)

type CatePage = response.PageData[Page]

type Page struct {
	ID            int       `json:"id"`
	AppCategoryID int       `json:"app_category_id"`
	Name          string    `json:"name"`
	IsActive      bool      `json:"is_active" default:"true"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
	// Categorry     datatypes.JSON `json:"categorry"`
	Form sharedutil.JsonArray `json:"form"`
}

type PageCreate struct {
	AppCategoryID int    `json:"app_category_id"`
	Name          string `json:"name"`
	IsActive      *bool  `json:"is_active"`
}
