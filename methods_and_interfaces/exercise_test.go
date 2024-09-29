package methodsandinterfaces

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_IpStringer(t *testing.T) {
	item := IPAddr{127, 0, 0, 1}
	assert.Equal(t, "127.0.0.1", item.String())
}

func Test_handleError(t *testing.T) {
	_, result := Sqrt(-10)
	assert.Equal(t, result.Error(), "cannot Sqrt negitive number : -10")
}
