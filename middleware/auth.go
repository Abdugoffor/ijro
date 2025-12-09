package middleware

import (
	"ijro-nazorat/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		// Header
		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"error": "missing token",
			})
		}

		// Format: Bearer TOKEN
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"error": "invalid token format",
			})
		}

		tokenString := parts[1]

		// Parse
		claims, err := helper.ParseJWT(tokenString)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"error": "invalid or expired token",
			})
		}

		// Token valid → claims ni contextga yozamiz
		ctx.Set("user", claims)

		return next(ctx)
	}
}
