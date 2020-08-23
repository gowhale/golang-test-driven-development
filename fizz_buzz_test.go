package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeIsFizz(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(3), "FIZZ")
}

func TestFiveIsBuzz(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(5), "BUZZ")
}

func TestFifteenIsFizzBuzz(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(15), "FIZZBUZZ")
}

func TestOneIsOne(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(1), "1")
}