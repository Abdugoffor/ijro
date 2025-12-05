package user_cmd

import (
	user_handler "ijro-nazorat/modul/user/hadnler"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(route *echo.Echo, db *gorm.DB, log *log.Logger) {
	routerGroup := route.Group("/admin")
	{
		user_handler.NewUserHandler(routerGroup, db, log)
	}
}
