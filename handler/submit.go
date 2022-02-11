package handler

import (
	"errors"
	"github.com/codigician/submit/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type (
	SubmissionReq struct {
		Content string `json:"content"`
		Lang    string `json:"lang"`
	}

	SubmissionRsp struct {
		ExecutionMs time.Duration `json:"execution_ms"`
		Status      bool          `json:"status"`
	}

	SubmissionsRsp []SubmissionRsp
)

type SubmitService interface {
	Submit(qID string, s service.Solution) ([]service.Submission, error)
}

type Submit struct {
	service SubmitService
}

func NewSubmit(submitService SubmitService) *Submit {
	return &Submit{submitService}
}

func (s *Submit) RegisterRoutes(router *echo.Echo) {
	router.POST("/submit/:questionID", s.Submit)
}

func (s *Submit) Submit(c echo.Context) error {
	questionID := c.Param("questionID")

	var req SubmissionReq
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if err := req.validate(); err != nil {
		return echo.ErrBadRequest
	}

	rsp, err := s.service.Submit(questionID, req.toSolution())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, fromSubmission(rsp))
}

func (s *SubmissionReq) validate() error {
	if s.Content == "" || s.Lang == "" {
		return errors.New("content or language cannot be empty")
	}
	return nil
}

func (s *SubmissionReq) toSolution() service.Solution {
	return service.Solution{
		Content: s.Content,
		Lang:    s.Lang,
	}
}

func fromSubmission(s []service.Submission) SubmissionsRsp {
	var rsp SubmissionsRsp
	for _, submission := range s {
		rsp = append(rsp, SubmissionRsp{
			ExecutionMs: submission.ExecutionMs,
			Status:      submission.Status,
		})
	}
	return rsp
}
