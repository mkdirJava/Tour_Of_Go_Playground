package typesstructcsslicesmaps

import "fmt"

// Pointers are variables that will point to a memory location

func Pointer() *string {
	var stringPointer *string
	stringValue := "one two three"
	// The line below sets the pointer stringPointer to the memory location
	// of stringValue
	stringPointer = &stringValue
	// you can get the value of the item the pointer is pointing at via
	// *stringPointer, see test
	return stringPointer
}

func PointersAllowNilsToBeReturned() *string {
	return nil
}

// like in Java Arrays are limited in their size
func ArrayActionInit() []int {
	// There are many ways to init a array, one example
	myArray := []int{1, 2, 3, 4}
	return myArray
}

// Slices are dynamically sized, like the ArrayList in Java
func Slices() []int {
	items := []int{1, 2, 3, 4}
	// The [start:end]
	// [:2] from the start to the second element
	// [2:] from the second element to the end
	// [:] make a reference to it all
	return items[2:3]
}

// Slices operate like pointers of the array.
// If the array changes then the slices also change
func SliceChange() []int {
	items := []int{1, 2, 3, 4}
	returnningItem := items[2:]
	items[3] = 999
	return returnningItem
}

// The Length of a slice is from the reference it was created with
// whereas the capcacity is the size of the underlying array
// see test for useage
func SliceLengthAndCapacity() [20]int {
	return [20]int{1, 2, 3, 4}
}

func MakingThings() []int {
	return make([]int, 5)
}

func AppendingToSlice() []int {
	items := []int{1, 2, 3, 4}
	items = append(items, 5)
	return items
}

func loopingOverSlices() {
	items := []int{1, 2, 3, 4, 5}
	// You can do _ for variables you dont use
	for index, value := range items {
		println(fmt.Sprintf("the Index is %d and the value %v", index, value))
	}
}

func InitMap() map[string]string {
	items := map[string]string{"home": "home", "away": "away"}
	return items
}

func ReturnAValueFromAFunc(arg string) string {
	myFunctionReference := func(arg string) string {
		return fmt.Sprintf("Hi there %s", arg)
	}
	return myFunctionReference(arg)
}
