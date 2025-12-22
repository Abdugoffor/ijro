package category_handler

import (
	"fmt"
	category_dto "ijro-nazorat/modul/category/dto"
	category_service "ijro-nazorat/modul/category/service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"git.sriss.uz/shared/shared_service/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type categoryHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service category_service.CategoryService
}

func NewCategoryHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) categoryHandler {
	handler := categoryHandler{
		db:      db,
		log:     log,
		service: category_service.NewCategoryService(db),
	}

	routes := gorm.Group("/category")
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

func (handler *categoryHandler) All(ctx echo.Context) error {
	req := request.Request(ctx)
	var query category_dto.Filter
	{
		if err := ctx.Bind(&query); err != nil {
			return err
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		switch query.Status {
		case "open":
			tx = tx.Where("deleted_at IS NULL")
		case "deleted":
			tx = tx.Unscoped().Where("deleted_at IS NOT NULL")
		default:
			tx = tx.Unscoped()
		}

		if query.Name != "" {
			name := "%" + strings.ToLower(query.Name) + "%"
			tx = tx.Where("LOWER(categories.name) LIKE ?", name)
		}

		sortColumn := "categories.id ASC"

		if query.Column != "" && query.Sort != "" {
			sortColumn = fmt.Sprintf("categories.%s %s", query.Column, query.Sort)
		}

		tx = tx.Group("categories.id").Order(sortColumn)

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

func (handler *categoryHandler) Show(ctx echo.Context) error {
	req := request.Request(ctx)

	idParam := ctx.Param("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 64)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if parsedID > 0 {
			tx = tx.Where("id = ?", parsedID)
		}
		return tx
	}

	data, err := handler.service.Show(req.Context(), filter)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *categoryHandler) Create(ctx echo.Context) error {
	var req category_dto.CreateOrUpdate
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

func (handler *categoryHandler) Update(ctx echo.Context) error {
	idParam := ctx.Param("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 64)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
		}
	}

	var req category_dto.CreateOrUpdate
	{
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if parsedID > 0 {
			tx = tx.Where("id = ?", parsedID)
		}
		return tx
	}

	data, err := handler.service.Update(ctx, filter, req)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *categoryHandler) Delete(ctx echo.Context) error {
	idParam := ctx.Param("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 64)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if parsedID > 0 {
			tx = tx.Where("id = ?", parsedID)
		}
		return tx
	}

	err = handler.service.Delete(ctx, filter)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, echo.Map{"message": "success delete data"})
}

func (handler *categoryHandler) Restore(ctx echo.Context) error {
	req := request.Request(ctx)
	idParam := ctx.Param("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 64)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if parsedID > 0 {
			tx = tx.Where("id = ?", parsedID)
		}
		return tx
	}

	data, err := handler.service.Restore(req.Context(), filter)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *categoryHandler) ForceDelete(ctx echo.Context) error {
	idParam := ctx.Param("id")
	parsedID, err := strconv.ParseUint(idParam, 10, 64)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if parsedID > 0 {
			tx = tx.Where("id = ?", parsedID)
		}
		return tx
	}

	err = handler.service.ForceDelete(ctx, filter)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, echo.Map{"message": "success delete data"})
}

func GetID(ctx echo.Context) (int64, error) {
	idParam := ctx.Param("id")

	parsedID, err := strconv.ParseInt(idParam, 10, 64)
	{
		if err != nil {
			return 0, ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
		}
	}

	return parsedID, nil
}
