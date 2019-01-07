package fizzbuzz_test

import (
	. "fizzbuzz/fizzbuzz"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzInput3ShouldHaveNothing(t *testing.T) {
	assert.Equal(t, true, IsFizz(3))
}

func TestFizzInput5ShouldHaveBuzz(t *testing.T) {
	assert.Equal(t, false, IsFizz(5))
}
