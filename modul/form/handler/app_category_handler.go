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

type AppCategoryHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service form_service.AppCategoryService
}

func NewAppCategoryHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := AppCategoryHandler{
		db:      db,
		log:     log,
		service: form_service.NewAppCategoryService(db),
	}

	routes := gorm.Group("/app-category")
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

func (handler *AppCategoryHandler) All(ctx echo.Context) error {
	req := request.Request(ctx)

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Select(`
        app_category.id,
        app_category.name,
        app_category.is_active,
        app_category.created_at,
        app_category.updated_at,
        app_category.deleted_at,
        COALESCE(
            jsonb_agg(
                DISTINCT jsonb_build_object(
                    'id', page.id,
                    'name', page.name,
                    'is_active', page.is_active,
                    'created_at', page.created_at,
                    'updated_at', page.updated_at,
                    'deleted_at', page.deleted_at,
                    'form', (
                        SELECT COALESCE(
                            jsonb_agg(
                                jsonb_build_object(
                                    'id', f.id,
                                    'name', f.name,
                                    'label', f.label,
                                    'is_required', f.is_required,
                                    'is_active', f.is_active,
                                    'created_at', f.created_at,
                                    'updated_at', f.updated_at,
                                    'deleted_at', f.deleted_at
                                )
                            ),
                            '[]'::jsonb
                        )
                        FROM form f
                        WHERE f.page_id = page.id
                    )
                )
            ) FILTER (WHERE page.id IS NOT NULL),
            '[]'::jsonb
        ) AS pages
    `).Joins(`
        LEFT JOIN page ON app_category.id = page.app_category_id
    `).Group("app_category.id")
	}

	data, err := handler.service.All(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, data)
}

func (handler *AppCategoryHandler) Show(ctx echo.Context) error {
	return nil
}

func (handler *AppCategoryHandler) Create(ctx echo.Context) error {

	var req form_dto.AppCreateOrUpdate
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

func (handler *AppCategoryHandler) Update(ctx echo.Context) error {
	return nil
}

func (handler *AppCategoryHandler) Delete(ctx echo.Context) error {
	return nil
}

func (handler *AppCategoryHandler) Restore(ctx echo.Context) error {
	return nil
}

func (handler *AppCategoryHandler) ForceDelete(ctx echo.Context) error {
	return nil
}
