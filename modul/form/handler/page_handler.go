package form_handler

import (
	form_dto "ijro-nazorat/modul/form/dto"
	form_service "ijro-nazorat/modul/form/service"
	"log"
	"net/http"

	"git.sriss.uz/shared/shared_service/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PageHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service form_service.PageService
}

func NewPageHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := PageHandler{
		db:      db,
		log:     log,
		service: form_service.NewPageService(db),
	}

	routes := gorm.Group("/page")
	{
		routes.GET("", handler.All)
		routes.GET("/:id", handler.Show)
		routes.POST("", handler.Create)
		routes.PUT("/:id", handler.Update)
		routes.DELETE("/:id", handler.Delete)
		routes.PATCH("/restore/:id", handler.Restore)
		routes.DELETE("/force/:id", handler.ForceDelete)
	}
}

func (handler *PageHandler) All(ctx echo.Context) error {
	req := request.Request(ctx)

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Select(`
		page.id,
		page.app_category_id,
		page.name,
		page.is_active,
		page.created_at,
		page.updated_at,
		page.deleted_at,

		COALESCE(
			json_agg(
				json_build_object(
					'id', categorry.id,
					'name', categorry.name
				)
			) FILTER (WHERE categorry.id IS NOT NULL),
			'[]'
		) AS categorry,

		COALESCE(
			json_agg(
				json_build_object(
					'id', form.id,
					'name', form.name,
					'label', form.label,
					'is_active', form.is_active,
					'is_required', form.is_required
				)
			) FILTER (WHERE form.id IS NOT NULL),
			'[]'
		) AS form
	`).
			Joins("LEFT JOIN app_category AS categorry ON page.app_category_id = categorry.id").
			Joins("LEFT JOIN form ON page.id = form.page_id").
			Group("page.id")
	}

	data, err := handler.service.All(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, data)
}

func (handler *PageHandler) Show(ctx echo.Context) error {
	return nil
}

func (handler *PageHandler) Create(ctx echo.Context) error {
	var req form_dto.PageCreate
	{
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	data, err := handler.service.Create(ctx, req)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *PageHandler) Update(ctx echo.Context) error {
	return nil
}

func (handler *PageHandler) Delete(ctx echo.Context) error {
	return nil
}

func (handler *PageHandler) Restore(ctx echo.Context) error {
	return nil
}

func (handler *PageHandler) ForceDelete(ctx echo.Context) error {
	return nil
}
