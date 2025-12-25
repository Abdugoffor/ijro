package shop_handler

import (
	"ijro-nazorat/helper"
	shop_dto "ijro-nazorat/modul/shop/dto"
	shop_service "ijro-nazorat/modul/shop/service"
	"log"
	"net/http"
	"strconv"

	"git.sriss.uz/shared/shared_service/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type shopHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service shop_service.ShopService
}

func NewShopHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) shopHandler {
	handler := shopHandler{
		db:      db,
		log:     log,
		service: shop_service.NewShopService(db),
	}
	routes := gorm.Group("/shop")
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

func (handler *shopHandler) All(ctx echo.Context) error {
	req := request.Request(ctx)

	filter := func(tx *gorm.DB) *gorm.DB {
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

func (handler *shopHandler) Show(ctx echo.Context) error {
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

func (handler *shopHandler) Create(ctx echo.Context) error {
	var req shop_dto.ShopChange
	{
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}
	req.Slug = helper.Slug(req.Name)

	data, err := handler.service.Create(ctx, req)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *shopHandler) Update(ctx echo.Context) error {
	var req shop_dto.ShopChange
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

	return ctx.JSON(http.StatusOK, data)
}

func (handler *shopHandler) Delete(ctx echo.Context) error {
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

func (handler *shopHandler) Restore(ctx echo.Context) error {
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

func (handler *shopHandler) ForceDelete(ctx echo.Context) error {
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
