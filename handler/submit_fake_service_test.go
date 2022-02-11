package handler_test

import (
	"errors"
	"github.com/codigician/submit/service"
)

type FakeSubmitService struct {
}

func (f *FakeSubmitService) Submit(qID string, s service.Solution) ([]service.Submission, error) {
	if qID == "1" {
		return nil, errors.New("abc")
	}

	return nil, nil
}
