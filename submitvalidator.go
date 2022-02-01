package main

import (
	"errors"
)

type RequestValidator struct {
}

func (v *RequestValidator) Validate(req SubmitRequest) error {
	if !supportedLanguage(req.Lang) {
		return errors.New("language is not supported")
	}

	return nil
}

func supportedLanguage(lang string) bool {
	return supportedLanguages[lang]
}
