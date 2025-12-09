package user_handler

import (
	"fmt"
	user_dto "ijro-nazorat/modul/user/dto"
	user_service "ijro-nazorat/modul/user/service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service user_service.UserService
}

func NewUserHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) userHandler {
	handler := userHandler{
		db:      db,
		log:     log,
		service: user_service.NewUserService(db),
	}

	routes := gorm.Group("/user")
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

func (handler *userHandler) All(ctx echo.Context) error {
	var query user_dto.Filter
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
			tx = tx.Where("LOWER(users.name) LIKE ?", name)
		}

		if query.Email != "" {
			email := "%" + strings.ToLower(query.Email) + "%"
			tx = tx.Where("LOWER(users.email) LIKE ?", email)
		}

		if query.Role != "" {
			tx = tx.Where("users.role = ?", query.Role)
		}

		sortColumn := "users.id ASC"

		if query.Column != "" && query.Sort != "" {
			sortColumn = fmt.Sprintf("users.%s %s", query.Column, query.Sort)
		}

		return tx.Preload("Country").Group("users.id").Order(sortColumn)
	}

	res, err := handler.service.All(ctx, filter)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, res)
}

func (handler *userHandler) Show(ctx echo.Context) error {
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
		return tx.Preload("Country")
	}

	data, err := handler.service.Show(ctx, filter)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *userHandler) Create(ctx echo.Context) error {
	var req user_dto.CreateOrUpdate
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

func (handler *userHandler) Update(ctx echo.Context) error {
	var req user_dto.CreateOrUpdate
	{
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
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

func (handler *userHandler) Delete(ctx echo.Context) error {
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

func (handler *userHandler) Restore(ctx echo.Context) error {
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

	data, err := handler.service.Restore(ctx, filter)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *userHandler) ForceDelete(ctx echo.Context) error {
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
