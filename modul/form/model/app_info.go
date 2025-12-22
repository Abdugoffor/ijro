package form_model

import "time"

type AppInfo struct {
	ID        int       `json:"id"`
	AppID     int       `json:"app_id"`
	PageID    int       `json:"page_id"`
	FormID    int       `json:"form_id"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (AppInfo) TableName() string {
	return "app_info"
}
