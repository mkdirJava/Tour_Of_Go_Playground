package mysecondimplementation

import (
	"testing"

	i "mkdirjava/golangSite/methods_and_interfaces/my_first_interface"

	"github.com/go-playground/assert/v2"
)

// Notice here the we need to get a pointer if of the variable if we
// want to use the pointer reference of the object_orientated
// we want to use the pointer reciever as it allows modification of internal state
var unit i.MyFirstInterface = &MySecondImplementation{}

func Test_MyFirstImplemation(t *testing.T) {
	unit.SetString("bye")
	assert.Equal(t, "bye", unit.GetExportedString())
}
