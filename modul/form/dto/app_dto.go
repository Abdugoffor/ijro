package form_dto

type ApplicationCreate struct {
	AppCategoryID int `json:"app_category_id"`
	PageID        int `json:"page_id"`
	Answers       []AppAnsware
}

type AppAnsware struct {
	FormID int    `json:"form_id"`
	Answer string `json:"answer"`
}
