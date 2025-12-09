package middleware

import (
	"ijro-nazorat/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			claims, ok := ctx.Get("user").(*helper.Claims)
			if !ok {
				return ctx.JSON(http.StatusUnauthorized, echo.Map{
					"error": "unauthorized",
				})
			}

			if claims.Role != role {
				return ctx.JSON(http.StatusForbidden, echo.Map{
					"error": "access denied",
				})
			}

			return next(ctx)
		}
	}
}
