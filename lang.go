package main

type SupportedLanguages map[string]bool

var (
	supportedLanguages = SupportedLanguages{
		"golang":     true,
		"python3":    true,
		"javascript": true,
		"nodejs":     true,
		"c":          true,
		"c++":        true,
		"java8":      true,
	}
)
