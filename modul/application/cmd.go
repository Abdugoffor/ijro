package application_cmd

import (
	"ijro-nazorat/middleware"
	application_handler "ijro-nazorat/modul/application/handler"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(route *echo.Echo, db *gorm.DB, log *log.Logger) {
	routerGroup := route.Group("/admin", middleware.JWTMiddleware, middleware.RoleMiddleware("admin"))
	{
		application_handler.NewApplicationHandler(routerGroup, db, log)
	}
}
