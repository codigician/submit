package service

import "time"

type Solution struct {
	Content string
	Lang    string
}

type Submission struct {
	ExecutionMs time.Duration
	Status      bool
}
