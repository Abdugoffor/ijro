package form_model

import "time"

type Form struct {
	ID         int       `json:"id"`
	PageID     int       `json:"page_id"`
	Name       string    `json:"name"`
	Label      string    `json:"label"`
	IsRequired bool      `json:"is_required" default:"false"`
	IsActive   bool      `json:"is_active" default:"true"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

func (Form) TableName() string {
	return "form"
}
