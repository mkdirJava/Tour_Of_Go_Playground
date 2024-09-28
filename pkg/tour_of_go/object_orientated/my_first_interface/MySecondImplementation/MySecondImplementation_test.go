package mysecondimplementation

import (
	"testing"

	i "mkdirjava/golangSite/pkg/tour_of_go/object_orientated/my_first_interface"

	"github.com/go-playground/assert/v2"
)

var unit i.MyFirstInterface = MySecondImplementation{}

func Test_MyFirstImplemation(t *testing.T) {
	assert.Equal(t, "bye", unit.GetExportedString())
}
