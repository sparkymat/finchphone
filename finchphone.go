package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	e.Logger.Panic(e.Start(":8080"))
}
