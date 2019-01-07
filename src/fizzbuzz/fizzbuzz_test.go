// +build integration

package fizzbuzz_test

import (
	. "fizzbuzz/fizzbuzz"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz2ShouldBe2(t *testing.T) {
	assert.Equal(t, "2", FizzBuzz("2"))
}

func TestFizzBuzz3ShouldBeFizz(t *testing.T) {
	assert.Equal(t, "Fizz", FizzBuzz("3"))
}

func TestFizzBuzz5ShouldBeBuzz(t *testing.T) {
	assert.Equal(t, "Buzz", FizzBuzz("5"))
}

func TestFizzBuzz15ShouldBeFizzBuzz(t *testing.T) {
	assert.Equal(t, "FizzBuzz", FizzBuzz("15"))
}
