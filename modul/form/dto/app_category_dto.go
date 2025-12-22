package form_dto

import (
	"git.sriss.uz/shared/shared_service/response"
	"gorm.io/datatypes"
)

type AppCategoryPage = response.PageData[AppResponse]

type AppResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	IsActive  bool           `json:"is_active"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt string         `json:"deleted_at"`
	Pages     datatypes.JSON `json:"pages"`
	// Pages     sharedutil.JsonArray `json:"pages"`
}

type AppCreateOrUpdate struct {
	Name     string `json:"name"`
	IsActive *bool  `json:"is_active"`
}
