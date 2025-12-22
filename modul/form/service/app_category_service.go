package form_service

import (
	"context"
	"ijro-nazorat/helper"
	form_dto "ijro-nazorat/modul/form/dto"
	form_model "ijro-nazorat/modul/form/model"

	"git.sriss.uz/shared/shared_service/pg"
	"git.sriss.uz/shared/shared_service/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AppCategoryService interface {
	All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*form_dto.AppCategoryPage, error)
	Create(ctx echo.Context, req form_dto.AppCreateOrUpdate) (form_dto.AppResponse, error)
	Update(ctx context.Context, filter pg.Filter, req form_dto.AppCreateOrUpdate) (form_dto.AppResponse, error)
	Delete(ctx context.Context, filter pg.Filter) error
	Restore(ctx context.Context, filter pg.Filter) (form_dto.AppResponse, error)
	ForceDelete(ctx context.Context, filter pg.Filter) error
}

type appCategoryService struct {
	db *gorm.DB
}

func NewAppCategoryService(db *gorm.DB) AppCategoryService {
	return &appCategoryService{
		db: db,
	}
}

func (service *appCategoryService) All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*form_dto.AppCategoryPage, error) {
	return pg.PageWithScan[form_model.AppCategory, form_dto.AppResponse](service.db, paginate, filter)
}

func (service *appCategoryService) Create(ctx echo.Context, req form_dto.AppCreateOrUpdate) (form_dto.AppResponse, error) {
	model := form_model.AppCategory{
		Name: req.Name,
	}

	if req.IsActive != nil {
		model.IsActive = *req.IsActive
	} else {
		model.IsActive = true
	}

	if err := service.db.Create(&model).Error; err != nil {
		return form_dto.AppResponse{}, err
	}

	return form_dto.AppResponse{
		ID:        model.ID,
		Name:      model.Name,
		IsActive:  model.IsActive,
		CreatedAt: helper.FormatDate(model.CreatedAt),
		UpdatedAt: helper.FormatDate(model.UpdatedAt),
		DeletedAt: helper.FormatDate(model.UpdatedAt),
	}, nil
}

func (service *appCategoryService) Update(ctx context.Context, filter pg.Filter, req form_dto.AppCreateOrUpdate) (form_dto.AppResponse, error) {
	return form_dto.AppResponse{}, nil
}

func (service *appCategoryService) Delete(ctx context.Context, filter pg.Filter) error {
	return nil
}

func (service *appCategoryService) Restore(ctx context.Context, filter pg.Filter) (form_dto.AppResponse, error) {
	return form_dto.AppResponse{}, nil
}

func (service *appCategoryService) ForceDelete(ctx context.Context, filter pg.Filter) error {
	return nil
}
