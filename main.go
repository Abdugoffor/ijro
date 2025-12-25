package main

import (
	"ijro-nazorat/config"
	"ijro-nazorat/helper"
	category_cmd "ijro-nazorat/modul/category"
	shop_cmd "ijro-nazorat/modul/shop"
	"ijro-nazorat/seeder"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	helper.LoadEnv()
	config.DBConnect()
	seeder.DBSeeders()

	route := echo.New()

	// route.Renderer = views.NewRenderer()

	route.Validator = config.NewValidator()

	category_cmd.Cmd(route, config.DB, log.Default())
	shop_cmd.Cmd(route, config.DB, log.Default())

	route.Logger.Fatal(route.Start(":" + helper.ENV("HTTP_PORT")))
}
