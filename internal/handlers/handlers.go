package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/Lovodia/restapi/internal/models"
	"github.com/Lovodia/restapi/internal/storage"
	"github.com/Lovodia/restapi/internal/usecase"

	"github.com/labstack/echo/v4"
)

func PostHandler(logger *slog.Logger, store *storage.ResultStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		var nums models.Numbers
		if err := c.Bind(&nums); err != nil {
			logger.Error("Failed to bind request body", "error", err.Error())
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

		key := strconv.FormatInt(time.Now().UnixNano(), 10)
		store.Save(key, total)

		return c.JSON(http.StatusOK, resp)
	}
}
func MultiplyHandler(logger *slog.Logger, store *storage.ResultStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req models.Numbers
		if err := c.Bind(&req); err != nil {
			logger.Error("Failed to bind multiply request", "error", err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid data format")
		}
		if req.Values == nil {
			logger.Info("Received numbers", "values", "nil slice")
		} else {
			logger.Info("Received numbers", "values", req.Values)
		}

		multiply := usecase.CalculatedMultiply(req.Values)

		logger.Info("Calculated multiply", "multiply", multiply)
		resp := models.MultiplyResponse{Multiply: multiply}

		key := strconv.FormatInt(time.Now().UnixNano(), 10)
		store.Save(key, multiply)

		return c.JSON(http.StatusOK, resp)
	}
}

func GetAllResultsHandler(logger *slog.Logger, store *storage.ResultStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		results := store.GetAll()
		return c.JSON(http.StatusOK, results)
	}
}
