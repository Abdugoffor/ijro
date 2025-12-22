package auth_handler

import (
	"ijro-nazorat/helper"
	"ijro-nazorat/middleware"
	auth_dto "ijro-nazorat/modul/auth/dto"
	auth_service "ijro-nazorat/modul/auth/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type authHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service auth_service.AuthService
}

func NewAuthHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) authHandler {
	handler := authHandler{
		db:      db,
		log:     log,
		service: auth_service.NewAuthService(db),
	}

	routes := gorm.Group("/auth")
	{
		routes.POST("/login", handler.Login)
		routes.POST("/refresh-token", handler.RefreshToken)
		routes.POST("/logout", handler.Logout)
		routes.GET("/me", handler.Me, middleware.JWTMiddleware)
	}

	return handler
}

func (handler *authHandler) Login(ctx echo.Context) error {
	var req auth_dto.Login
	{
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		if err := ctx.Validate(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}

	data, err := handler.service.Login(req)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *authHandler) RefreshToken(ctx echo.Context) error {
	var req auth_dto.RefreshToken
	{
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		if err := ctx.Validate(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}

	data, err := handler.service.RefreshToken(req)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(200, data)
}

func (handler *authHandler) Logout(ctx echo.Context) error {
	return nil
}

func (handler *authHandler) Me(ctx echo.Context) error {
	// JWTMiddleware ichida yozilgan claims ni olish
	claims, ok := ctx.Get("user").(*helper.Claims)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}
	// Response qaytarish
	return ctx.JSON(http.StatusOK, echo.Map{
		"user_id":    claims.UserID,
		"country_id": claims.CountryId,
		"email":      claims.Email,
		"role":       claims.Role,
		"name":       claims.Name,
	})
}
