package fizzbuzz_test

import (
	. "fizzbuzz/fizzbuzz"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuzzInput3ShouldHaveNothing(t *testing.T) {
	assert.Equal(t, false, IsBuzz(3))
}

func TestBuzzInput5ShouldHaveBuzz(t *testing.T) {
	assert.Equal(t, true, IsBuzz(5))
}
