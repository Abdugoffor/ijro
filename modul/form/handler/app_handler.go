package form_handler

import (
	form_dto "ijro-nazorat/modul/form/dto"
	form_service "ijro-nazorat/modul/form/service"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AppHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service form_service.AppService
}

func NewAppHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := AppHandler{
		db:      db,
		log:     log,
		service: form_service.NewAppService(db),
	}

	routes := gorm.Group("/app")
	{
		routes.POST("", handler.Create)
	}
}

func (handler *AppHandler) Create(ctx echo.Context) error {
	var req form_dto.ApplicationCreate
	{
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(400, map[string]string{
				"error": err.Error(),
			})
		}
	}

	appID, err := handler.service.Create(ctx, req)
	{
		if err != nil {
			return ctx.JSON(500, map[string]string{
				"error": err.Error(),
			})
		}
	}

	return ctx.JSON(201, map[string]any{
		"app_id": appID,
	})
}
