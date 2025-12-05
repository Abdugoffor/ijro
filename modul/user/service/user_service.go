package user_service

import (
	"ijro-nazorat/helper"
	user_dto "ijro-nazorat/modul/user/dto"
	user_model "ijro-nazorat/modul/user/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserService interface {
	All(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (helper.PaginatedResponse[user_dto.Response], error)
	Show(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (user_dto.Response, error)
	Create(ctx echo.Context, req user_dto.CreateOrUpdate) (user_dto.Response, error)
	Update(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB, req user_dto.CreateOrUpdate) (user_dto.Response, error)
	Delete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error
	Restore(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (user_dto.Response, error)
	ForceDelete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (service *userService) All(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (helper.PaginatedResponse[user_dto.Response], error) {
	var models []user_model.User

	res, err := helper.Paginate(ctx, service.db.Scopes(filter), &models, 10)
	{
		if err != nil {
			return helper.PaginatedResponse[user_dto.Response]{}, err
		}
	}

	var data []user_dto.Response
	{
		for _, model := range models {
			data = append(data, user_dto.ToResponse(model))
		}
	}

	return helper.PaginatedResponse[user_dto.Response]{
		Data: data,
		Meta: res.Meta,
	}, nil

}
func (service *userService) Show(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (user_dto.Response, error) {
	var model user_model.User
	{
		if err := service.db.Scopes(filter).First(&model).Error; err != nil {
			return user_dto.Response{}, err
		}
	}

	res := user_dto.ToResponse(model)
	return res, nil
}

func (service *userService) Create(ctx echo.Context, req user_dto.CreateOrUpdate) (user_dto.Response, error) {
	model := user_model.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
		CountryID: req.CountryID,
	}

	if req.IsActive != nil {
		model.IsActive = *req.IsActive
	} else {
		model.IsActive = true
	}

	if err := service.db.Create(&model).Error; err != nil {
		return user_dto.Response{}, err
	}

	res := user_dto.ToResponse(model)
	return res, nil
}
func (service *userService) Update(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB, req user_dto.CreateOrUpdate) (user_dto.Response, error) {
	var model user_model.User
	{
		if err := service.db.Scopes(filter).First(&model).Error; err != nil {
			return user_dto.Response{}, err
		}
	}

	model.Name = req.Name
	model.Email = req.Email
	model.Password = req.Password
	model.Role = req.Role
	model.CountryID = req.CountryID
	model.IsActive = *req.IsActive

	if err := service.db.Save(&model).Error; err != nil {
		return user_dto.Response{}, err
	}

	res := user_dto.ToResponse(model)
	return res, nil
}
func (service *userService) Delete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error {
	var model user_model.User
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
func (service *userService) Restore(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) (user_dto.Response, error) {
	var model user_model.User
	{
		if err := service.db.Unscoped().Scopes(filter).First(&model).Error; err != nil {
			return user_dto.Response{}, err
		}
	}

	if err := service.db.Model(&model).Unscoped().Update("deleted_at", nil).Error; err != nil {
		return user_dto.Response{}, err
	}

	return user_dto.Response{}, nil
}
func (service *userService) ForceDelete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error {
	var model user_model.User
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
