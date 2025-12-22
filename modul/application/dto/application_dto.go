package application_dto

import (
	application_model "ijro-nazorat/modul/application/model"
	"time"
)

type Create struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	File        string `json:"file"`
	CategoryID  int    `json:"category_id"`
	CountryID   int    `json:"country_id"` // to'g'ri nom
}

type Update struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	File        string `json:"file"`
	Status      string `json:"status"`
	CategoryID  int    `json:"category_id"`
	CountryID   int    `json:"country_id"` // to'g'ri nom
}

type Response struct {
	ID int `json:"id"`
	// UserID      int    `json:"user_id"`
	User        string `json:"user"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	File        string `json:"file"`
	// CategoryID  int    `json:"category_id"`
	Category string `json:"category"`
	// CountryID   int    `json:"country_id"` // to'g'ri nom
	Country   string   `json:"country"`
	Status    string   `json:"status"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Answers   []Answer `json:"answers"`
}

type Filter struct {
	Name       string `json:"name" query:"name"`
	CountryID  *int   `json:"country_id" query:"country_id"`
	CategoryID *int   `json:"category_id" query:"category_id"`
	Status     string `json:"status" query:"status"`
	Sort       string `json:"sort" query:"sort"`
	Column     string `json:"column" query:"column"`
}

type StatusUpdate struct {
	Status string `json:"status"`
}

type AnswerCreate struct {
	ApplicationID uint   `json:"application_id"`
	Answer        string `json:"answer"`
}

type Answer struct {
	ID        int       `json:"id"`
	User      string    `json:"user"`
	Answer    string    `json:"answer"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
type AnswerResponse struct {
	ID          int       `json:"id"`
	User        string    `json:"user"`
	Application Response  `json:"application"`
	Answer      string    `json:"answer"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func ToResponse(application application_model.Application) Response {
	// DTO Answer ga conversion
	var answers []Answer
	for _, ans := range application.Answers {
		answers = append(answers, Answer{
			ID:        ans.ID,
			User:      ans.User.Name,
			Answer:    ans.Answer,
			Status:    ans.Status,
			CreatedAt: ans.CreatedAt,
		})
	}

	return Response{
		ID:          application.ID,
		User:        application.User.Name,
		Name:        application.Name,
		Description: application.Description,
		Image:       application.Image,
		File:        application.File,
		Category:    application.Category.Name,
		Country:     application.Country.Name,
		Status:      application.Status,
		CreatedAt:   application.CreatedAt,
		UpdatedAt:   application.UpdatedAt,
		Answers:     answers,
	}
}

func ToAnswerResponse(answer application_model.Answer) AnswerResponse {
	return AnswerResponse{
		ID:          answer.ID,
		User:        answer.User.Name,
		Application: ToResponse(answer.Application),
		Answer:      answer.Answer,
		Status:      answer.Status,
		CreatedAt:   answer.CreatedAt,
	}
}
