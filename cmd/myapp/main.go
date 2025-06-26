package main

import (
	"log"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Lovodia/-REST-API/internal/handlers"
	"github.com/Lovodia/-REST-API/pkg/config"
)

func main() {
	cfg := config.LoadConfig()

	logger := slog.New(slog.NewTextHandler(log.Writer()))

	e := echo.New()

	e.Logger.SetOutput(nil)

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	e.POST("/calculate-sum", handlers.PostHandler(logger))

	port := cfg.ServerPort
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
