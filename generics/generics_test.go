package generics

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_generics(t *testing.T) {
	unit := My_List[int]{
		value: []int{1, 2, 3},
	}
	assert.Equal(t, 1, unit.Find(1.0))
}
