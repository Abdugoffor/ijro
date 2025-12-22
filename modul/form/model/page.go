package form_model

import "time"

type Page struct {
	ID            int       `json:"id"`
	AppCategoryID int       `json:"app_category_id"`
	Name          string    `json:"name"`
	IsActive      bool      `json:"is_active" default:"true"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

func (Page) TableName() string {
	return "page"
}
