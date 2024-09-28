package myfirstinterface

// We define the struct and then hang the method that needs
// to match the interface method signatures
type MyFirstImplementation struct{}

func (m MyFirstImplementation) GetExportedString() string {
	return "hi"
}
