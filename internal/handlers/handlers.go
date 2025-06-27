package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Lovodia/restapi/internal/models"
	"github.com/Lovodia/restapi/internal/usecase"
	loggers "github.com/Lovodia/restapi/pkg/logger"

	"github.com/labstack/echo/v4"
)

func PostHandler(logger *slog.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var nums models.Numbers
		if err := c.Bind(&nums); err != nil {
			loggers.SafeLogError(logger, "Failed to bind request body", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid data format")
		}

		if nums.Values == nil {
			loggers.SafeLogInfo(logger, "Received numbers", "values", "nil slice")
		} else {
			loggers.SafeLogInfo(logger, "Received numbers", "values", nums.Values)
		}

		total := usecase.CalculateSum(nums.Values)

		resp := models.SumResponse{Sum: total}
		loggers.SafeLogInfo(logger, "Calculated sum", "sum", total)

		return c.JSON(http.StatusOK, resp)
	}
}
