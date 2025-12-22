package form_service

import (
	"context"
	form_dto "ijro-nazorat/modul/form/dto"
	form_model "ijro-nazorat/modul/form/model"

	"git.sriss.uz/shared/shared_service/pg"
	"git.sriss.uz/shared/shared_service/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type FormService interface {
	All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*form_dto.FormPage, error)
	Create(ctx echo.Context, req form_dto.FormCreate) (form_dto.Form, error)
}

type formService struct {
	db *gorm.DB
}

func NewFormService(db *gorm.DB) FormService {
	return &formService{
		db: db,
	}
}

func (service *formService) All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*form_dto.FormPage, error) {
	return pg.PageWithScan[form_model.Form, form_dto.Form](service.db, paginate, filter)
}

func (service *formService) Create(ctx echo.Context, req form_dto.FormCreate) (form_dto.Form, error) {
	model := form_model.Form{
		Name:   req.Name,
		Label:  req.Label,
		PageID: req.PageID,
	}

	if req.IsActive != nil {
		model.IsActive = *req.IsActive
	} else {
		model.IsActive = true
	}

	if req.IsRequired != nil {
		model.IsRequired = *req.IsRequired
	} else {
		model.IsRequired = false
	}

	if err := service.db.Create(&model).Error; err != nil {
		return form_dto.Form{}, err
	}

	return form_dto.Form{
		ID:         model.ID,
		Label:      model.Label,
		Name:       model.Name,
		IsRequired: model.IsRequired,
		IsActive:   model.IsActive,
		// CreatedAt:  model.CreatedAt,
		// UpdatedAt:  model.UpdatedAt,
		// DeletedAt:  model.UpdatedAt,
	}, nil

}
