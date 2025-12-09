package auth_cmd

import (
	auth_handler "ijro-nazorat/modul/auth/handler"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(route *echo.Echo, db *gorm.DB, log *log.Logger) {
	routerGroup := route.Group("/admin")
	{
		auth_handler.NewAuthHandler(routerGroup, db, log)
	}
}
