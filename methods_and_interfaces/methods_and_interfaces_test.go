package methodsandinterfaces

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_assertions(t *testing.T) {
	var itemOne interfaceOne = TypeOne{}
	_, castedOk := itemOne.(TypeTwo)
	assert.Equal(t, castedOk, false)
}
