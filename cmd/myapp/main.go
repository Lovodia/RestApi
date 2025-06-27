package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Lovodia/restapi/internal/handlers"
	"github.com/Lovodia/restapi/pkg/config"
	"github.com/Lovodia/restapi/pkg/logger1" // импортируем новый пакет
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := logger1.NewLogger(cfg.Logger.Level)

	e := echo.New()
	e.Logger.SetOutput(os.Stdout)

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	e.POST("/calculate-sum", handlers.PostHandler(logger))

	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
