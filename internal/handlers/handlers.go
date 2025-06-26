package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Lovodia/-REST-API/internal/models"
	"github.com/Lovodia/-REST-API/internal/usecase"

	"github.com/labstack/echo/v4"
)

func PostHandler(logger *slog.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var nums models.Numbers
		if err := c.Bind(&nums); err != nil {

			if err != nil {
				logger.Error("Failed to bind request body", "error", err.Error())
			} else {
				logger.Error("Failed to bind request body: unknown error")
			}
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid data format")
		}

		if nums.Values == nil {
			logger.Info("Received numbers", "values", "nil slice")
		} else {
			logger.Info("Received numbers", "values", nums.Values)
		}

		total := usecase.CalculateSum(nums.Values)

		resp := models.SumResponse{Sum: total}
		logger.Info("Calculated sum", "sum", total)

		return c.JSON(http.StatusOK, resp)
	}
}
