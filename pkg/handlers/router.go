package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

// BuildRouter sets up the routes for the API
func BuildRouter(r *echo.Echo) {
	r.Use(middleware.Recover())
	r.Use(middleware.Secure())
	r.Use(middleware.BodyLimit("1KB"))
	r.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogLatency:       true,
		LogMethod:        true,
		LogURI:           true,
		LogStatus:        true,
		LogError:         true,
		LogContentLength: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("uri", v.URI).
				Int("status", v.Status).
				Dur("latency", v.Latency).
				Str("content_length", v.ContentLength).
				Msg("request received")
			return nil
		},
	}))
	r.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		ErrorMessage: "request timed out",
		Timeout:      1 * time.Second,
	}))

	r.Validator = NewValidator()

	r.GET("/min", handleMin)
	r.GET("/max", handleMax)
	r.GET("/avg", handleAverage)
	r.GET("/median", handleMedian)
	r.GET("/percentile", handlePercentile)
}
