package middleware

import (
	"fmt"
	"net/http"

	"github.com/J-Thompson12/math-api/pkg/equations"
	"github.com/labstack/echo/v4"
)

type Params struct {
	List       []int `json:"list" validate:"required,min=1"`
	Quantifier int   `json:"quantifier"`
}

const keyEquation = "equation"

func InitEquation(c echo.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			param, err := newParams(c)
			if err != nil {
				return err
			}

			var e equations.Equation
			if param.Quantifier == 0 {
				e = equations.NewEquation(param.List)
			} else {
				e = equations.NewEquationWithQuantifier(param.List, param.Quantifier)
			}
			SetEquation(c, e)
			next(c)
			return nil
		}
	}
}

func SetEquation(c echo.Context, e equations.Equation) {
	c.Set(keyEquation, e)
}

func GetEquation(c echo.Context) equations.Equation {
	return c.Get(keyEquation).(equations.Equation)
}

func newParams(c echo.Context) (Params, error) {
	var req Params
	if err := c.Bind(&req); err != nil {
		return req, echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid request: %w ", err))
	}

	if err := c.Validate(req); err != nil {
		return Params{}, echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return req, nil
}
