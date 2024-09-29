package object_orientated

// Golangs answear to a class is a struct
// To hang method to a struct use a reciver method
// The struct is there to hold variables

type My_First_Exported_struct struct {
	MyExportedField     string
	myNoneExportedField string
}

// this is an example of a method implementation that will hang
// off a struct
func (m *My_First_Exported_struct) A_PublicMethod() string {
	return "Hi there world"
}

// Usually there is a factory method to create new structs to be used in
// a program, usually this is done as a exported Method

func NewMy_First_Exported_Struct(arg string) My_First_Exported_struct {
	return My_First_Exported_struct{
		myNoneExportedField: arg,
		MyExportedField:     arg,
	}
}
