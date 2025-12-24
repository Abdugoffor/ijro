package form_dto

import (
	"encoding/json"

	"git.sriss.uz/shared/shared_service/response"
)

type AppPage = response.PageData[AppInfo]

type AppInfo struct {
	ID       int             `json:"id"`
	Category json.RawMessage `json:"category"` // object → OK
	Pages    json.RawMessage `json:"pages"`    // array → JsonArray
	// Category  sharedutil.JsonObject `json:"category"` // object → OK
	// Pages     sharedutil.JsonArray  `json:"pages"`    // array → JsonArray
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"` // nullable
}
type AppInfoPage struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
