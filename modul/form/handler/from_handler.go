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

type FormHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service form_service.FormService
}

func NewFormHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := FormHandler{
		db:      db,
		log:     log,
		service: form_service.NewFormService(db),
	}

	routes := gorm.Group("/form")
	{
		routes.GET("", handler.All)
		routes.POST("", handler.Create)
	}
}

func (handler FormHandler) All(ctx echo.Context) error {
	req := request.Request(ctx)

	filter := func(tx *gorm.DB) *gorm.DB {

		tx = tx.Select(`
		form.id,
		form.label,
		form.name,
		form.is_active,
		form.created_at,
		form.updated_at,
		form.deleted_at,
		COALESCE(
			json_agg(
				json_build_object(
					'id', page.id,
					'name', page.name
				)
			) FILTER (WHERE page.id IS NOT NULL),
			'[]'
		) AS page
	`).
			Joins("LEFT JOIN page AS page ON form.page_id = page.id").
			Group("form.id")

		return tx
	}

	data, err := handler.service.All(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, data)
}

func (handler FormHandler) Create(ctx echo.Context) error {

	var req form_dto.FormCreate
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
