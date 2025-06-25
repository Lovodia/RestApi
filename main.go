package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Numbers struct {
	Values []float64 `json:"values"`
}

type SumResponse struct {
	Sum float64 `json:"sum"`
}

func PostHandler(c echo.Context) error {
	var nums Numbers
	if err := c.Bind(&nums); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid data format")
	}

	sum := 0.0
	for _, v := range nums.Values {
		sum += v
	}

	resp := SumResponse{Sum: sum}
	return c.JSON(http.StatusOK, resp)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/post", PostHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
