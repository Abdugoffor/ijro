package form_dto

import (
	"time"

	"git.sriss.uz/shared/shared_service/response"
)

type FormPage = response.PageData[Form]

type Form = struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Label      string `json:"label"`
	IsRequired bool   `json:"is_required"`
	IsActive   bool   `json:"is_active"`
	// Page       datatypes.JSON `json:"page"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type FormCreate struct {
	PageID     int    `json:"page_id"`
	Name       string `json:"name"`
	Label      string `json:"label"`
	IsRequired *bool  `json:"is_required"`
	IsActive   *bool  `json:"is_active"`
}
