package myfirstinterface

// Like in most languages that support interfaces
// the private and public methods are defined here
// and the implementing class must implement them

type MyFirstInterface interface {
	GetExportedString() string
	SetString(string)
}
