package handlers

import (
	"net/http"

	"log/slog"

	"github.com/Lovodia/-REST-API/internal/models"
	"github.com/Lovodia/-REST-API/internal/usecase"

	"github.com/labstack/echo/v4"
)

func PostHandler(logger *slog.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var nums models.Numbers
		if err := c.Bind(&nums); err != nil {
			logger.Error("Failed to bind request body", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid data format")
		}

		logger.Info("Received numbers", "values", nums.Values)

		total := usecase.CalculateSum(nums.Values)

		resp := models.SumResponse{Sum: total}
		logger.Info("Calculated sum", "sum", total)

		return c.JSON(http.StatusOK, resp)
	}
}
