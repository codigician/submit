package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	defer func() {
		assert.Equal(t, "not implemented", recover())
	}()

	main()
}
