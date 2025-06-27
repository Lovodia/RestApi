package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Lovodia/restapi/internal/handlers"
	"github.com/Lovodia/restapi/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Инициализация логгера с указанным уровнем
	var logLevel slog.Level
	switch cfg.Logger.Level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, opts))

	e := echo.New()
	e.Logger.SetOutput(os.Stdout)

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	e.POST("/calculate-sum", handlers.PostHandler(logger))

	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
