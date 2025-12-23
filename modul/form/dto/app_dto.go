package form_dto

import (
	"git.sriss.uz/shared/shared_service/response"
	"gorm.io/datatypes"
)

type AppPage = response.PageData[AppInfo]

type AppInfo struct {
	ID        int            `json:"id"`
	Category  datatypes.JSON `json:"category"`
	Pages     datatypes.JSON `json:"pages"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt *string        `json:"deleted_at,omitempty"` // nullable

}

type ApplicationCreate struct {
	AppCategoryID int `json:"app_category_id"`
	PageID        int `json:"page_id"`
	Answers       []AppAnsware
}

type AppAnsware struct {
	FormID int    `json:"form_id"`
	Answer string `json:"answer"`
}
