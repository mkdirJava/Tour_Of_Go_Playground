package package_example

import "fmt"

func noneExportedMethod(arg string) string {
	return fmt.Sprintf("I said hi from the none exported methodpackage %s", arg)
}

func ExportedMethod(arg string) string {
	return fmt.Sprintf("I said hi from the exported method %s", arg)
}
