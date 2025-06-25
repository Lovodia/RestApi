package handlers

import (
	"net/http"

	"github.com/Lovodia/-REST-API/internal/models"
	"github.com/Lovodia/-REST-API/internal/sum"

	"github.com/labstack/echo/v4"
)

func PostHandler(c echo.Context) error {
	var nums models.Numbers
	if err := c.Bind(&nums); err != nil {
		c.Logger().Errorf("Failed to bind request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid data format")
	}

	c.Logger().Infof("Received numbers: %v", nums.Values)

	total := sum.CalculateSum(nums.Values)

	resp := models.SumResponse{Sum: total}
	c.Logger().Infof("Calculated sum: %f", total)

	return c.JSON(http.StatusOK, resp)
}
