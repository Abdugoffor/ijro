package form_dto

import (
	"git.sriss.uz/shared/shared_service/response"
	"git.sriss.uz/shared/shared_service/sharedutil"
)

type AppCategoryPage = response.PageData[AppResponse]

type AppResponse struct {
	ID        int                  `json:"id"`
	Name      string               `json:"name"`
	IsActive  bool                 `json:"is_active"`
	CreatedAt string               `json:"created_at"`
	UpdatedAt string               `json:"updated_at"`
	DeletedAt string               `json:"deleted_at"`
	Pages     sharedutil.JsonArray `json:"pages"`
	// Pages     sharedutil.JsonArray `json:"pages"`
}

type AppCreateOrUpdate struct {
	Name     string `json:"name"`
	IsActive *bool  `json:"is_active"`
}
