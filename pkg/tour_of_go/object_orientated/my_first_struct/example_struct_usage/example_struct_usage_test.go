package object_orientated_useage

import (
	o "mkdirjava/golangSite/pkg/tour_of_go/object_orientated/my_first_struct"
	"testing"

	"github.com/go-playground/assert/v2"
)

var (
	message = "Hi there world"
	unit    = o.NewMy_First_Exported_Struct(message)
)

func Test_MyFirst_Exported_Struct_Method(t *testing.T) {
	result := unit.A_PublicMethod()
	assert.Equal(t, message, result)
}

// comment this code block in and watch it fail, the myNoneExportedField is
// Not exported and will fail compilation

// func Test_fails_compilation(t *testing.T){
// 	unit.myNoneExportedField
// }
