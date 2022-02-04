package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type SubmitRequest struct {
	Lang       string `json:"lang"`
	Content    string `json:"content"`
	QuestionID string `json:"question_id"`
}

type Result struct {
	Duration time.Duration `json:"duration"`
	Status   string        `json:"status"`
}

type SubmitService interface {
	Submit(ctx context.Context, req SubmitRequest) ([]Result, error)
}

type handler struct {
	service SubmitService
}

func (h *handler) HandleSubmit(c echo.Context) error {
	var reqBody SubmitRequest
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	var requestValidator RequestValidator
	if validationError := requestValidator.Validate(reqBody); validationError != nil {
		return c.JSON(http.StatusBadRequest, validationError.Error())
	}

	//c.Request().Header.Get("Authorization")
	res, err := h.service.Submit(c.Request().Context(), reqBody)
	if err != nil {
		return c.JSON(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
