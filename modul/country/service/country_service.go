package country_service

import (
	"ijro-nazorat/helper"
	country_dto "ijro-nazorat/modul/country/dto"
	country_model "ijro-nazorat/modul/country/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CountryService interface {
	All(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (helper.PaginatedResponse[country_dto.Response], error)
	Show(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (country_dto.Response, error)
	Create(ctx echo.Context, req country_dto.CreateOrUpdate) (country_dto.Response, error)
	Update(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB, req country_dto.CreateOrUpdate) (country_dto.Response, error)
	Delete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error
	Restore(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (country_dto.Response, error)
	ForceDelete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error
}

type countryService struct {
	db *gorm.DB
}

func NewCountryService(db *gorm.DB) CountryService {
	return &countryService{db: db}
}

func (service *countryService) All(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (helper.PaginatedResponse[country_dto.Response], error) {
	var models []country_model.Country

	res, err := helper.Paginate(ctx, service.db.Scopes(filter), &models, 10)
	{
		if err != nil {
			return helper.PaginatedResponse[country_dto.Response]{}, err
		}
	}

	var data []country_dto.Response
	{
		for _, model := range models {
			data = append(data, country_dto.ToResponse(model))
		}
	}

	return helper.PaginatedResponse[country_dto.Response]{
		Data: data,
		Meta: res.Meta,
	}, nil
}

func (service *countryService) Show(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (country_dto.Response, error) {
	var model country_model.Country
	{
		if err := service.db.Scopes(filter).First(&model).Error; err != nil {
			return country_dto.Response{}, err
		}
	}

	res := country_dto.ToResponse(model)
	return res, nil
}

func (service *countryService) Create(ctx echo.Context, req country_dto.CreateOrUpdate) (country_dto.Response, error) {

	model := country_model.Country{
		Name: req.Name,
	}

	if req.IsActive != nil {
		model.IsActive = *req.IsActive
	} else {
		model.IsActive = true
	}

	if err := service.db.Create(&model).Error; err != nil {
		return country_dto.Response{}, err
	}
	res := country_dto.ToResponse(model)
	return res, nil
}

func (service *countryService) Update(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB, req country_dto.CreateOrUpdate) (country_dto.Response, error) {
	var model country_model.Country
	{
		if err := service.db.Scopes(filter).First(&model).Error; err != nil {
			return country_dto.Response{}, err
		}
	}

	model.Name = req.Name
	model.IsActive = *req.IsActive
	{
		if err := service.db.Save(&model).Error; err != nil {
			return country_dto.Response{}, err
		}
	}

	res := country_dto.ToResponse(model)

	return res, nil
}

func (service *countryService) Delete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error {
	var model country_model.Country
	{
		if err := service.db.Scopes(filter).First(&model).Error; err != nil {
			return err
		}
	}

	if err := service.db.Delete(&model).Error; err != nil {
		return err
	}

	return nil
}

func (service *countryService) Restore(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (country_dto.Response, error) {
	var model country_model.Country
	{
		if err := service.db.Unscoped().Scopes(filter).First(&model).Error; err != nil {
			return country_dto.Response{}, err
		}
	}

	if err := service.db.Model(&model).Update("deleted_at", nil).Error; err != nil {
		return country_dto.Response{}, err
	}

	res := country_dto.ToResponse(model)

	return res, nil
}

func (service *countryService) ForceDelete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error {
	var model country_model.Country
	{
		if err := service.db.Unscoped().Scopes(filter).First(&model).Error; err != nil {
			return err
		}
	}

	if err := service.db.Unscoped().Delete(&model).Error; err != nil {
		return err
	}

	return nil
}
