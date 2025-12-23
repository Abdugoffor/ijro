package main

import (
	"ijro-nazorat/config"
	"ijro-nazorat/helper"
	application_cmd "ijro-nazorat/modul/application"
	auth_cmd "ijro-nazorat/modul/auth"
	category_cmd "ijro-nazorat/modul/category"
	country_cmd "ijro-nazorat/modul/country"
	form_cmd "ijro-nazorat/modul/form"
	user_cmd "ijro-nazorat/modul/user"
	"ijro-nazorat/seeder"
	"ijro-nazorat/views"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	helper.LoadEnv()
	config.DBConnect()
	seeder.DBSeeders()

	route := echo.New()

	route.Renderer = views.NewRenderer()

	route.Validator = config.NewValidator()

	auth_cmd.Cmd(route, config.DB, log.Default())
	category_cmd.Cmd(route, config.DB, log.Default())
	country_cmd.Cmd(route, config.DB, log.Default())
	user_cmd.Cmd(route, config.DB, log.Default())
	application_cmd.Cmd(route, config.DB, log.Default())

	form_cmd.Cmd(route, config.DB, log.Default())

	route.Logger.Fatal(route.Start(":" + helper.ENV("HTTP_PORT")))
}
