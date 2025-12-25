package shop_cmd

import (
	shop_handler "ijro-nazorat/modul/shop/handler"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(route *echo.Echo, db *gorm.DB, log *log.Logger) {
	routerGroup := route.Group("/admin")
	{
		shop_handler.NewShopHandler(routerGroup, db, log)
	}
}
