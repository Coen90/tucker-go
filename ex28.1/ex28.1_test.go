package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	assert.Equal(t, 81, square(9), "square(9) should be 81")

	assert.Equal(t, 9, square(3), "should be 9")
}