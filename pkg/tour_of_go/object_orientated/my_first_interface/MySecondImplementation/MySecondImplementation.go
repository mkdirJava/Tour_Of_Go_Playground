package mysecondimplementation

type MySecondImplementation struct{}

func (m MySecondImplementation) GetExportedString() string {
	return "bye"
}

func (m MySecondImplementation) getNoneExportedString() string {
	return "bye"
}
