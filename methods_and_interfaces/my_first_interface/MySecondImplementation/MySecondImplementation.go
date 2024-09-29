package mysecondimplementation

type MySecondImplementation struct {
	message string
}

func (m *MySecondImplementation) GetExportedString() string {
	return m.message
}

func (m *MySecondImplementation) getNoneExportedString() string {
	return m.message
}

func (m *MySecondImplementation) SetString(arg string) {
	m.message = arg
}
