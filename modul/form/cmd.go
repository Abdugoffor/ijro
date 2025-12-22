package form_cmd

import (
	form_handler "ijro-nazorat/modul/form/handler"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(route *echo.Echo, db *gorm.DB, log *log.Logger) {
	routerGroup := route.Group("/admin")
	{
		form_handler.NewAppCategoryHandler(routerGroup, db, log)
		form_handler.NewPageHandler(routerGroup, db, log)
		form_handler.NewFormHandler(routerGroup, db, log)
		form_handler.NewAppHandler(routerGroup, db, log)
	}
}
