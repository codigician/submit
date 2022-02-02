package main

import (
	"errors"
)

type RequestValidator struct {
}

func (v *RequestValidator) Validate(req SubmitRequest) error {
	if req.Lang == "" {
		return errors.New("Parameter cannot be null")
	}

	if req.QuestionID == "" {
		return errors.New("Parameter cannot be null")
	}

	if req.Content == "" {
		return errors.New("Parameter cannot be null")
	}

	return nil
}
