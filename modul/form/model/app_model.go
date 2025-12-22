package form_model

import "time"

type App struct {
	ID            int       `json:"id"`
	AppCategoryID int       `json:"app_category_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

func (App) TableName() string {
	return "app"
}
