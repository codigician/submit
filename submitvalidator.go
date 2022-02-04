package main

import (
	"errors"
)

type RequestValidator struct {
}

func (v *RequestValidator) Validate(req SubmitRequest) error {
	if req.Lang == "" {
		return errors.New("Lang is required")
	}

	if req.QuestionID == "" {
		return errors.New("question_id is required")
	}

	if req.Content == "" {
		return errors.New("Content is required")
	}

	return nil
}
