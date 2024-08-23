package main

import (
	"github.com/J-Thompson12/math-api/pkg/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	handlers.BuildRouter(e)
	e.Logger.Fatal(e.Start(":8000"))
}
