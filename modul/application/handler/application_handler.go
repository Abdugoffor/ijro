package application_handler

import (
	"ijro-nazorat/middleware"
	application_dto "ijro-nazorat/modul/application/dto"
	application_service "ijro-nazorat/modul/application/service"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type applicationHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service application_service.ApplicationService
}

func NewApplicationHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) applicationHandler {
	handler := applicationHandler{
		db:      db,
		log:     log,
		service: application_service.NewApplicationService(db),
	}

	routes := gorm.Group("/application", middleware.JWTMiddleware, middleware.RoleMiddleware("admin"))
	{
		routes.GET("", handler.All)
		routes.GET("/:id", handler.Show)
		routes.POST("", handler.Create)
		routes.PUT("/:id", handler.Update)
		routes.DELETE("/:id", handler.Delete)
		routes.PATCH("/restore/:id", handler.Restore)
		routes.DELETE("/force/:id", handler.ForceDelete)
	}

	return handler
}

func (handler *applicationHandler) All(ctx echo.Context) error {
	var query application_dto.Filter
	{
		if err := ctx.Bind(&query); err != nil {
			return err
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		if query.Name != "" {
			search := "%" + strings.ToLower(query.Name) + "%"
			tx = tx.Where("LOWER(applications.name) LIKE ?", search)
		}

		if query.CountryID != nil {
			tx = tx.Where("country_id = ?", *query.CountryID)
		}

		if query.CategoryID != nil {
			tx = tx.Where("category_id = ?", *query.CategoryID)
		}

		if query.Status != "" {
			tx = tx.Where("LOWER(applications.status) = ?", strings.ToLower(query.Status))
		}

		return tx.
			Preload("User").
			Preload("Category").
			Preload("Country").
			Preload("Answers")
	}

	res, err := handler.service.All(ctx, filter)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, res)
}

func (handler *applicationHandler) Show(ctx echo.Context) error {
	return nil
}

func (handler *applicationHandler) Create(ctx echo.Context) error {
	var req application_dto.Create
	{
		if err := ctx.Bind(&req); err != nil {
			return err
		}
	}

	data, err := handler.service.Create(ctx, req)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, data)
}

func (handler *applicationHandler) Update(ctx echo.Context) error {
	return nil
}

func (handler *applicationHandler) Delete(ctx echo.Context) error {
	return nil
}

func (handler *applicationHandler) Restore(ctx echo.Context) error {
	return nil
}

func (handler *applicationHandler) ForceDelete(ctx echo.Context) error {
	return nil
}
