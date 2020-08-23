package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeIsFizz(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(3), "FIZZ")
}
