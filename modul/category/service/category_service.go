package category_service

import (
	"context"
	category_dto "ijro-nazorat/modul/category/dto"
	category_model "ijro-nazorat/modul/category/model"

	"git.sriss.uz/shared/shared_service/pg"
	"git.sriss.uz/shared/shared_service/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryService interface {
	All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*category_dto.CategoryPage, error)
	Show(ctx context.Context, filter pg.Filter) (*category_dto.Response, error)
	Create(ctx echo.Context, req category_dto.CreateOrUpdate) (category_dto.Response, error)
	Update(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB, category category_dto.CreateOrUpdate) (category_dto.Response, error)
	Delete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error
	Restore(ctx context.Context, filter pg.Filter) (category_dto.Response, error)
	ForceDelete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error
}

type categoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) CategoryService {
	return &categoryService{db: db}
}

func (service *categoryService) All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*category_dto.CategoryPage, error) {
	return pg.PageWithScan[category_model.Category, category_dto.Response](service.db, paginate, filter)
}

func (service *categoryService) Show(ctx context.Context, filter pg.Filter) (*category_dto.Response, error) {
	return pg.FindOneWithScan[category_model.Category, category_dto.Response](service.db, filter)
}

func (service *categoryService) Create(ctx echo.Context, req category_dto.CreateOrUpdate) (category_dto.Response, error) {
	model := category_model.Category{
		Name: req.Name,
	}

	if req.IsActive != nil {
		model.IsActive = *req.IsActive
	} else {
		model.IsActive = true
	}

	if err := service.db.Create(&model).Error; err != nil {
		return category_dto.Response{}, err
	}

	res := category_dto.Response{
		ID:        model.ID,
		Name:      model.Name,
		IsActive:  model.IsActive,
		CreatedAt: model.Name,
		UpdatedAt: model.Name,
		DeletedAt: model.Name,
	}

	return res, nil
}

func (service *categoryService) Update(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB, category category_dto.CreateOrUpdate) (category_dto.Response, error) {
	var model category_model.Category
	{
		if err := service.db.Scopes(filter).First(&model).Error; err != nil {
			return category_dto.Response{}, err
		}
	}

	model.Name = category.Name
	model.IsActive = *category.IsActive
	{
		if err := service.db.Save(&model).Error; err != nil {
			return category_dto.Response{}, err
		}
	}

	res := category_dto.Response{
		ID:        model.ID,
		Name:      model.Name,
		IsActive:  model.IsActive,
		CreatedAt: model.Name,
		UpdatedAt: model.Name,
		DeletedAt: model.Name,
	}

	return res, nil

}

func (service *categoryService) Delete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error {
	var model category_model.Category
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

func (service *categoryService) Restore(ctx context.Context, filter func(tx *gorm.DB) *gorm.DB) (category_dto.Response, error) {
	var model category_model.Category
	{
		if err := service.db.Unscoped().Scopes(filter).First(&model).Error; err != nil {
			return category_dto.Response{}, err
		}
	}

	if err := service.db.Model(&model).Unscoped().Update("deleted_at", nil).Error; err != nil {
		return category_dto.Response{}, err
	}
	return category_dto.Response{}, nil
}

func (service *categoryService) ForceDelete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error {
	var model category_model.Category
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
