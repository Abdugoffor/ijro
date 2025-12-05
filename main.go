package main

import (
	"ijro-nazorat/config"
	"ijro-nazorat/helper"
	category_cmd "ijro-nazorat/modul/category"
	country_cmd "ijro-nazorat/modul/country"
	user_cmd "ijro-nazorat/modul/user"
	"ijro-nazorat/seeder"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	helper.LoadEnv()

	config.DBConnect()

	seeder.DBSeeders()

	route := echo.New()

	category_cmd.Cmd(route, config.DB, log.Default())
	country_cmd.Cmd(route, config.DB, log.Default())
	user_cmd.Cmd(route, config.DB, log.Default())

	route.Logger.Fatal(route.Start(":" + helper.ENV("HTTP_PORT")))
}
