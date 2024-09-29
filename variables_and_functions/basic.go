package basics

import "strconv"

// This is how you define a function / method

func AnExportedMethod() {
	println("I can be used anywhere")
}

func aNoneExportedMethod() {
	println("I can only be used in the same package")
}

func IReturnMoreThanOneThing() (string, string) {
	return "one", "two"
}

// Notice that the return type has variable names,
// they are used in the method do define their return value
func NamedResults() (a string, b string) {
	a = "one"
	b = "two"
	return
}

func ShortVariableDeclaration() string {
	message := "hi"
	return message
}

// These variables are inited, they dont have to be
// they will take default values if they are not inited.
// The variable first letter of the names are uppercase
// therefore are exported
// notice that these are constonts and cannot be reassigned,
// like final in Java
const One, Two, Three string = "one", "two", "three"

// There converstion functions outthere that look like Static methods (If you come from Java)
// eg in this strconv.FormatInt
// Failing these use the naming type.(variable_To_Cast/Convert)
// Rememeber once casted the reference variable is of that Type
// and will be treated so by the program
func TypeConverstion(arg int64) (string, float64) {
	convertedString := strconv.FormatInt(arg, 10)
	convertedFloat := float64(arg)

	return convertedString, convertedFloat
}
