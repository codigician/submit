package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	h := handler{
		service: &Service{},
	}
	e.POST("/submit", h.HandleSubmit)
	e.Logger.Fatal(e.Start(":8000"))
}
