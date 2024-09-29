package controlflow

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_squareRoot(t *testing.T) {
	assert.Equal(t, int32(10), int32(SquareRoot(100.00)))
}
