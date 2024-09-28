package myfirstinterface

import (
	"testing"

	i "mkdirjava/golangSite/pkg/tour_of_go/object_orientated/my_first_interface"

	"github.com/go-playground/assert/v2"
)

// The interest is here,
// unit is the interface reference, MyFirstInterface
// the code will only recognise the mthods of the interface
// and subject to the package be available to use determined by
// exported reference and useage, ie is the usage inside the package
// in this case yes

var unit i.MyFirstInterface = MyFirstImplementation{}

func Test_MyFirstImplemation(t *testing.T) {
	assert.Equal(t, "hi", unit.GetExportedString())
}
