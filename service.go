package main

import (
	"context"
	"time"
)

type Service struct {
}

func (s *Service) Submit(ctx context.Context, req SubmitRequest) ([]Result, error) {

	//content := req.Content
	//language := req.Lang
	//questionId := req.QuestionID

	// dummy return value
	return []Result{
		{
			Duration: 5 * time.Second,
			Status:   "success",
		},
		{
			Duration: 2 * time.Second,
			Status:   "fail",
		},
	}, nil
}
