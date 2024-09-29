package myfirstinterface

// We define the struct and then hang the method that needs
// to match the interface method signatures
type MyFirstImplementation struct {
	message string
}

func (m MyFirstImplementation) GetExportedString() string {
	return m.message
}

func (m MyFirstImplementation) SetString(arg string) {
	m.message = arg
}
