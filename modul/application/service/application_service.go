package application_service

import (
	"errors"
	"ijro-nazorat/helper"
	application_dto "ijro-nazorat/modul/application/dto"
	application_model "ijro-nazorat/modul/application/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ApplicationService interface {
	All(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (helper.PaginatedResponse[application_dto.Response], error)
	Show(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (application_dto.Response, error)
	Create(ctx echo.Context, req application_dto.Create) (application_dto.Response, error)
	Update(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB, req application_dto.Update) (application_dto.Response, error)
	Delete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error
	Restore(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (application_dto.Response, error)
	ForceDelete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error
}

type applicationService struct {
	db *gorm.DB
}

func NewApplicationService(db *gorm.DB) ApplicationService {
	return &applicationService{db: db}
}

func (service *applicationService) All(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (helper.PaginatedResponse[application_dto.Response], error) {
	var models []application_model.Application
	res, err := helper.Paginate(ctx, service.db.Scopes(filter), &models, 10)
	{
		if err != nil {
			return helper.PaginatedResponse[application_dto.Response]{}, err
		}
	}

	var data []application_dto.Response
	{
		for _, model := range models {
			data = append(data, application_dto.ToResponse(model))
		}
	}

	return helper.PaginatedResponse[application_dto.Response]{
		Data: data,
		Meta: res.Meta,
	}, nil
}

func (service *applicationService) Show(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (application_dto.Response, error) {
	return application_dto.Response{}, nil
}

func (service *applicationService) Create(ctx echo.Context, req application_dto.Create) (application_dto.Response, error) {
	claims, ok := ctx.Get("user").(*helper.Claims)
	if !ok {
		return application_dto.Response{}, errors.New("unauthorized")
	}

	model := application_model.Application{
		Name:        req.Name,
		UserID:      claims.UserID,
		Description: req.Description,
		Image:       req.Image,
		File:        req.File,
		CategoryID:  req.CategoryID,
		CountryID:   req.CountryID, // to'g'ri nom
		Status:      "pending",
	}

	if err := service.db.Create(&model).Error; err != nil {
		return application_dto.Response{}, err
	}

	return application_dto.ToResponse(model), nil
}

func (service *applicationService) Update(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB, req application_dto.Update) (application_dto.Response, error) {
	return application_dto.Response{}, nil
}

func (service *applicationService) Delete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error {
	return nil
}

func (service *applicationService) Restore(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (application_dto.Response, error) {
	return application_dto.Response{}, nil
}

func (service *applicationService) ForceDelete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error {
	return nil
}
