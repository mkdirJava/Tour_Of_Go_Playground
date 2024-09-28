package package_example

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

var testMessage = "Bob"

func Test_ExportedTest(t *testing.T) {
	result := ExportedMethod(testMessage)
	assert.Equal(t, fmt.Sprintf("I said hi from the exported method %s", testMessage), result)
}

func Test_NoneExportedMethod(t *testing.T) {
	result := noneExportedMethod(testMessage)
	assert.Equal(t, fmt.Sprintf("I said hi from the none exported methodpackage %s", testMessage), result)
}
