package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Lovodia/-REST-API/internal/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/post", handlers.PostHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

//fghjk
