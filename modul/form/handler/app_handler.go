package form_handler

import (
	form_dto "ijro-nazorat/modul/form/dto"
	form_service "ijro-nazorat/modul/form/service"
	"log"

	"git.sriss.uz/shared/shared_service/request"
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
		routes.GET("", handler.All)
		routes.POST("", handler.Create)
	}
}

func (handler *AppHandler) All(ctx echo.Context) error {
	req := request.Request(ctx)

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Select(`
		app.id,
		to_char(app.created_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS created_at,
		to_char(app.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS updated_at,
		to_char(app.deleted_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS deleted_at,

		-- CATEGORY
		jsonb_build_object(
			'id', app_category.id,
			'name', app_category.name,
			'is_active', app_category.is_active
		) AS category,
				-- PAGES
				COALESCE(
					jsonb_agg(
						DISTINCT jsonb_build_object(
							'id', page.id,
							'name', page.name,
							'is_active', page.is_active,
							'forms', (
								SELECT COALESCE(
									jsonb_agg(
										jsonb_build_object(
											'id', form.id,
											'label', form.label,
											'name', form.name,
											'answer', app_info.answer
										)
									),
									'[]'::jsonb
								)
								FROM form
								LEFT JOIN app_info
									ON app_info.form_id = form.id
									AND app_info.app_id = app.id        -- ✅ FAQAT SHU ARIZA
								WHERE form.page_id = page.id           -- ✅ FAQAT SHU PAGE
							)
						)
					) FILTER (WHERE page.id IS NOT NULL),
					'[]'::jsonb
				) AS pages
			`).
			Joins("JOIN app_category ON app_category.id = app.app_category_id").
			Joins("LEFT JOIN page ON page.app_category_id = app_category.id").
			Group("app.id, app_category.id").
			Order("app.id DESC")
	}

	data, err := handler.service.All(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, data)
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
