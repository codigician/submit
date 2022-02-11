package main

import (
	"fmt"
	"github.com/codigician/submit/handler"
	"github.com/codigician/submit/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	submitHandler := handler.NewSubmit(&MockSubmitService{})
	submitHandler.RegisterRoutes(e)
	e.Start(":8080")
	fmt.Printf("hello world")
}

type MockSubmitService struct{}

func (f *MockSubmitService) Submit(qID string, s service.Solution) ([]service.Submission, error) {

	return nil, nil
}
