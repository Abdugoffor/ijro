package country_handler

import (
	"fmt"
	country_dto "ijro-nazorat/modul/country/dto"
	country_service "ijro-nazorat/modul/country/service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type contryHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service country_service.CountryService
}

func NewContryHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) contryHandler {
	handler := contryHandler{
		db:      db,
		log:     log,
		service: country_service.NewCountryService(db),
	}

	routes := gorm.Group("/country")
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

func (handler *contryHandler) All(ctx echo.Context) error {
	var query country_dto.Filter
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
			tx = tx.Where("LOWER(countries.name) LIKE ?", name)
		}

		sortColumn := "countries.created_at ASC"

		if query.Column != "" && query.Sort != "" {
			sortColumn = fmt.Sprintf("countries.%s %s", query.Column, query.Sort)
		}

		tx = tx.Group("countries.id").Order(sortColumn)

		return tx
	}

	data, err := handler.service.All(ctx, filter)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, data)
}

func (handler *contryHandler) Show(ctx echo.Context) error {
	idParam := ctx.Param("id")

	parsedID, err := strconv.ParseInt(idParam, 10, 64)
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

	data, err := handler.service.Show(ctx, filter)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *contryHandler) Create(ctx echo.Context) error {
	var req country_dto.CreateOrUpdate
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

func (handler *contryHandler) Update(ctx echo.Context) error {
	idParam := ctx.Param("id")

	parsedID, err := strconv.ParseInt(idParam, 10, 64)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
		}
	}

	var req country_dto.CreateOrUpdate
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

func (handler *contryHandler) Delete(ctx echo.Context) error {
	idParam := ctx.Param("id")

	parsedID, err := strconv.ParseInt(idParam, 10, 64)
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

func (handler *contryHandler) Restore(ctx echo.Context) error {
	idParam := ctx.Param("id")

	parsedID, err := strconv.ParseInt(idParam, 10, 64)
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

	data, err := handler.service.Restore(ctx, filter)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *contryHandler) ForceDelete(ctx echo.Context) error {

	idParam := ctx.Param("id")

	parsedID, err := strconv.ParseInt(idParam, 10, 64)
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
