package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=internal/view

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sparkymat/finchphone/internal/config"
	"github.com/sparkymat/finchphone/internal/database"
	"github.com/sparkymat/finchphone/internal/dbx"
	"github.com/sparkymat/finchphone/internal/route"
)

func main() {
	e := echo.New()

	appConfig, err := config.New()
	if err != nil {
		panic(err)
	}

	dbDriver, err := database.New(appConfig.DatabaseURL())
	if err != nil {
		log.Error(err)
		panic(err)
	}

	if err = dbDriver.AutoMigrate(); err != nil {
		log.Error(err)
		panic(err)
	}

	db := dbx.New(dbDriver.DB())

	route.Setup(e, appConfig, db)

	e.Logger.Panic(e.Start(":8080"))
}
