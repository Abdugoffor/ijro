package shop_service

import (
	"context"
	shop_dto "ijro-nazorat/modul/shop/dto"
	shop_model "ijro-nazorat/modul/shop/model"

	"git.sriss.uz/shared/shared_service/pg"
	"git.sriss.uz/shared/shared_service/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ShopService interface {
	All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*shop_dto.ShopPage, error)
	Show(ctx context.Context, filter pg.Filter) (*shop_dto.ShopResponse, error)
	Create(ctx echo.Context, req shop_dto.ShopChange) (shop_dto.ShopResponse, error)
	Update(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB, category shop_dto.ShopChange) (shop_dto.ShopResponse, error)
	Delete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error
	Restore(ctx context.Context, filter func(tx *gorm.DB) *gorm.DB) (shop_dto.ShopResponse, error)
	ForceDelete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error
}

type shopService struct {
	db *gorm.DB
}

func NewShopService(db *gorm.DB) ShopService {
	return &shopService{db: db}
}

func (service *shopService) All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*shop_dto.ShopPage, error) {
	return pg.PageWithScan[shop_model.Shop, shop_dto.ShopResponse](service.db, paginate, filter)
}

func (service *shopService) Show(ctx context.Context, filter pg.Filter) (*shop_dto.ShopResponse, error) {
	return pg.FindOneWithScan[shop_model.Shop, shop_dto.ShopResponse](service.db, filter)
}

func (service *shopService) Create(ctx echo.Context, req shop_dto.ShopChange) (shop_dto.ShopResponse, error) {
	model := shop_model.Shop{
		Name: req.Name,
		Slug: req.Slug,
	}

	if req.IsActive != nil {
		model.IsActive = *req.IsActive
	} else {
		model.IsActive = true
	}

	if err := service.db.Create(&model).Error; err != nil {
		return shop_dto.ShopResponse{}, err
	}

	res := shop_dto.ShopResponse{
		ID:        model.ID,
		Name:      model.Name,
		Slug:      model.Slug,
		IsActive:  model.IsActive,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}

	return res, nil
}

func (service *shopService) Update(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB, category shop_dto.ShopChange) (shop_dto.ShopResponse, error) {
	var model shop_model.Shop
	{
		if err := service.db.Scopes(filter).First(&model).Error; err != nil {
			return shop_dto.ShopResponse{}, err
		}
	}

	model.Name = category.Name
	model.IsActive = *category.IsActive
	{
		if err := service.db.Save(&model).Error; err != nil {
			return shop_dto.ShopResponse{}, err
		}
	}

	res := shop_dto.ShopResponse{
		ID:        model.ID,
		Name:      model.Name,
		Slug:      model.Slug,
		IsActive:  model.IsActive,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}

	return res, nil
}

func (service *shopService) Delete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error {
	err := pg.Delete[shop_model.Shop](service.db, nil, filter)
	{
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(404, "shop not found")
			}
			return echo.NewHTTPError(500, err.Error())
		}
	}
	return ctx.NoContent(204)
}

func (service *shopService) Restore(ctx context.Context, filter func(tx *gorm.DB) *gorm.DB) (shop_dto.ShopResponse, error) {
	var model shop_model.Shop
	{
		if err := service.db.Unscoped().Scopes(filter).First(&model).Error; err != nil {
			return shop_dto.ShopResponse{}, err
		}
	}

	if err := service.db.Model(&model).Update("deleted_at", nil).Error; err != nil {
		return shop_dto.ShopResponse{}, err
	}

	res := shop_dto.ShopResponse{
		ID:        model.ID,
		Name:      model.Name,
		IsActive:  model.IsActive,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}

	return res, nil
}

func (service *shopService) ForceDelete(ctx echo.Context, filter func(tx *gorm.DB) *gorm.DB) error {
	var model shop_model.Shop
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
