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

func TestThirtyIsFizzBuzz(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(30), "FIZZBUZZ")
}

func TestOneIsOne(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(1), "1")
}

func TestTwoIsTwo(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(2), "2")
}

func TestSixIsFizz(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(6), "FIZZ")
}

func TestTenIsBuzz(t *testing.T) {
	assert.EqualValues(t, FizzBuzz(10), "BUZZ")
}