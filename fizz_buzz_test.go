package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(3), "FIZZ")
}
