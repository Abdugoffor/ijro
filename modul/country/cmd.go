package country_cmd

import (
	"ijro-nazorat/middleware"
	country_handler "ijro-nazorat/modul/country/handler"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(route *echo.Echo, db *gorm.DB, log *log.Logger) {
	routerGroup := route.Group("/admin", middleware.JWTMiddleware, middleware.RoleMiddleware("admin"))
	{
		country_handler.NewContryHandler(routerGroup, db, log)
	}
}
