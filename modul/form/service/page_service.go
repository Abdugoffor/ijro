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

type PageService interface {
	All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*form_dto.CatePage, error)
	Show(ctx echo.Context, filter pg.Filter) (form_dto.Page, error)
	Create(ctx echo.Context, req form_dto.PageCreate) (form_dto.Page, error)
	Update(ctx echo.Context, filter pg.Filter, req form_dto.PageCreate) (form_dto.Page, error)
	Delete(ctx echo.Context, filter pg.Filter) error
	Restore(ctx echo.Context, filter pg.Filter) (form_dto.Page, error)
	ForceDelete(ctx echo.Context, filter pg.Filter) error
}

type pageService struct {
	db *gorm.DB
}

func NewPageService(db *gorm.DB) PageService {
	return &pageService{
		db: db,
	}
}

func (service *pageService) All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*form_dto.CatePage, error) {
	return pg.PageWithScan[form_model.Page, form_dto.Page](service.db, paginate, filter)
}

func (p *pageService) Show(ctx echo.Context, filter pg.Filter) (form_dto.Page, error) {
	return form_dto.Page{}, nil
}

func (service *pageService) Create(ctx echo.Context, req form_dto.PageCreate) (form_dto.Page, error) {
	model := form_model.Page{
		Name:          req.Name,
		AppCategoryID: req.AppCategoryID,
	}

	if req.IsActive != nil {
		model.IsActive = *req.IsActive
	} else {
		model.IsActive = true
	}

	if err := service.db.Create(&model).Error; err != nil {
		return form_dto.Page{}, err
	}

	return form_dto.Page{
		ID:            model.ID,
		AppCategoryID: model.AppCategoryID,
		Name:          model.Name,
		IsActive:      model.IsActive,
		// CreatedAt:     model.CreatedAt,
		// UpdatedAt:     model.UpdatedAt,
		// DeletedAt:     model.UpdatedAt,
	}, nil
}

func (p *pageService) Update(ctx echo.Context, filter pg.Filter, req form_dto.PageCreate) (form_dto.Page, error) {
	return form_dto.Page{}, nil
}

func (p *pageService) Delete(ctx echo.Context, filter pg.Filter) error {
	return nil
}

func (p *pageService) Restore(ctx echo.Context, filter pg.Filter) (form_dto.Page, error) {
	return form_dto.Page{}, nil
}

func (p *pageService) ForceDelete(ctx echo.Context, filter pg.Filter) error {
	return nil
}
