package application_handler

import (
	"ijro-nazorat/helper"
	"ijro-nazorat/middleware"
	application_dto "ijro-nazorat/modul/application/dto"
	application_service "ijro-nazorat/modul/application/service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type applicationClinetHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service application_service.ApplicationService
}

func NewApplicationClinetHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) applicationClinetHandler {
	handler := applicationClinetHandler{
		db:      db,
		log:     log,
		service: application_service.NewApplicationService(db),
	}

	routes := gorm.Group("/application-client", middleware.JWTMiddleware, middleware.RoleMiddleware("user"))
	{
		routes.GET("", handler.All)
		routes.GET("/:id", handler.Show)
		routes.POST("/create", handler.Create)
		routes.PUT("/:id", handler.Update)
	}

	return handler
}

func (handler *applicationClinetHandler) All(ctx echo.Context) error {
	claims, ok := ctx.Get("user").(*helper.Claims)
	{
		if !ok {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"error": "unauthorized",
			})
		}
	}

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

		if query.CategoryID != nil {
			tx = tx.Where("category_id = ?", *query.CategoryID)
		}

		if query.Status != "" {
			tx = tx.Where("LOWER(applications.status) = ?", strings.ToLower(query.Status))
		}

		tx = tx.Where("country_id = ?", claims.CountryId)

		return tx.
			Preload("User").
			Preload("Category").
			Preload("Country")
	}

	res, err := handler.service.All(ctx, filter)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, res)
}

func (handler *applicationClinetHandler) Show(ctx echo.Context) error {
	claims, ok := ctx.Get("user").(*helper.Claims)
	{
		if !ok {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"error": "unauthorized",
			})
		}
	}

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
		tx = tx.Where("country_id = ?", claims.CountryId)
		return tx.
			Preload("User").
			Preload("Category").
			Preload("Country")
	}

	res, err := handler.service.Show(ctx, filter)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, res)
}

func (handler *applicationClinetHandler) Create(ctx echo.Context) error {
	claims, ok := ctx.Get("user").(*helper.Claims)
	{
		if !ok {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"error": "unauthorized",
			})
		}
	}

	var req application_dto.AnswerCreate
	{
		if err := ctx.Bind(&req); err != nil {
			return err
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {

		if req.ApplicationID > 0 {
			tx = tx.Where("id = ?", req.ApplicationID)
		}
		tx = tx.Where("country_id = ?", claims.CountryId)
		return tx.
			Preload("User").
			Preload("Category").
			Preload("Country")
	}

	res, err := handler.service.AnswerCreate(ctx, filter, req)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, res)
}

func (handler *applicationClinetHandler) Update(ctx echo.Context) error {
	claims, ok := ctx.Get("user").(*helper.Claims)
	{
		if !ok {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"error": "unauthorized",
			})
		}
	}

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
		tx = tx.Where("country_id = ?", claims.CountryId)
		return tx.
			Preload("User").
			Preload("Category").
			Preload("Country")
	}

	req := application_dto.StatusUpdate{
		Status: "show",
	}

	res, err := handler.service.StatusUpdate(ctx, filter, req)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, res)
}
