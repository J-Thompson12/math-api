package handlers

import (
	"net/http"

	"github.com/J-Thompson12/math-api/pkg/equations"
	"github.com/J-Thompson12/math-api/pkg/middleware"
	"github.com/labstack/echo/v4"
)

// handleMin handles the /min endpoint
func handleMin(c echo.Context) error {
	e := middleware.GetEquation(c)
	result, err := e.Min()
	switch err.(type) {
	case nil:
		return c.JSON(http.StatusOK, result)
	case equations.ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, err)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err) // for this api this will never be hit but I added it because on a normal api it would be used
	}
}

// handleMax handles the /max endpoint
func handleMax(c echo.Context) error {
	e := middleware.GetEquation(c)
	result, err := e.Max()
	switch err.(type) {
	case nil:
		return c.JSON(http.StatusOK, result)
	case equations.ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, err)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
}

// handleAverage handles the /average endpoint
func handleAverage(c echo.Context) error {
	e := middleware.GetEquation(c)
	result, err := e.Average()
	switch err.(type) {
	case nil:
		return c.JSON(http.StatusOK, result)
	case equations.ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, err)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
}

// handleMedian handles the /median endpoint
func handleMedian(c echo.Context) error {
	e := middleware.GetEquation(c)
	result, err := e.Median()
	switch err.(type) {
	case nil:
		return c.JSON(http.StatusOK, result)
	case equations.ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, err)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
}

// handlePercentile handles the /percentile endpoint
func handlePercentile(c echo.Context) error {
	e := middleware.GetEquation(c)
	result, err := e.Percentile()
	switch err.(type) {
	case nil:
		return c.JSON(http.StatusOK, result)
	case equations.ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, err)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
}
