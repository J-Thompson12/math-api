package handlers

import (
	"fmt"
	"net/http"

	"github.com/J-Thompson12/math-api/pkg/equations"
	"github.com/labstack/echo/v4"
)

type Request struct {
	List       []int `json:"list" validate:"required,min=1"`
	Quantifier int   `json:"quantifier"`
}

// handleMin handles the /min endpoint
func handleMin(c echo.Context) error {
	req, err := newRequest(c)
	if err != nil {
		return err
	}

	result, err := equations.Min(req.List, req.Quantifier)
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
	req, err := newRequest(c)
	if err != nil {
		return err
	}

	result, err := equations.Max(req.List, req.Quantifier)
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
	req, err := newRequest(c)
	if err != nil {
		return err
	}

	result, err := equations.Average(req.List)
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
	req, err := newRequest(c)
	if err != nil {
		return err
	}

	result, err := equations.Median(req.List)
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
	req, err := newRequest(c)
	if err != nil {
		return err
	}

	result, err := equations.Percentile(req.List, req.Quantifier)
	switch err.(type) {
	case nil:
		return c.JSON(http.StatusOK, result)
	case equations.ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, err)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
}

// newRequest creates a new request from the echo context
func newRequest(c echo.Context) (Request, error) {
	var req Request
	if err := c.Bind(&req); err != nil {
		return req, echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid request: %w ", err))
	}

	if err := c.Validate(req); err != nil {
		return Request{}, echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return req, nil
}
