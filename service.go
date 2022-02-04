package main

import (
	"context"
	"errors"
	"time"
)

type Service struct {
}

func (s *Service) Submit(ctx context.Context, req SubmitRequest) ([]Result, error) {
	if !supportedLanguage(req.Lang) {
		return nil, errors.New("Language is not supported")
	}

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

func supportedLanguage(lang string) bool {
	return supportedLanguages[lang]
}
