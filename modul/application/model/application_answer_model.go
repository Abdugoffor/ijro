package application_model

import (
	user_model "ijro-nazorat/modul/user/model"
	"time"
)

type Answer struct {
	ID            int             `json:"id"`
	UserId        int             `json:"user_id"`
	User          user_model.User `json:"user" gorm:"foreignKey:UserId"`
	ApplicationId int             `json:"application_id"`
	Application   Application     `json:"application" gorm:"foreignKey:ApplicationId"`
	Answer        string          `json:"answer"`
	Status        string          `json:"status" gorm:"default:pending"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

func (Answer) TableName() string {
	return "application_answer"
}
