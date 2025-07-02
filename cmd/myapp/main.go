package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Lovodia/restapi/internal/handlers"
	"github.com/Lovodia/restapi/pkg/config"
	loggerSwitch "github.com/Lovodia/restapi/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {

		tempLogger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
		tempLogger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger := loggerSwitch.NewLogger(cfg.Logger.Level)

	e := echo.New()
	e.Logger.SetOutput(os.Stdout)

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	e.POST("/calculate-sum", handlers.PostHandler(logger))

	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
