package controlflow

import (
	"errors"
	"strconv"
)

func AFullCForLoop() (returnValue string) {
	returnValue = ""
	for counter := 0; counter <= 10; counter++ {
		returnValue = returnValue + strconv.FormatInt(int64(counter), 10)
	}
	return
}

// This form of the for loop can be considered a while loop,
// the colons are not needed here
func APartiaCForLoop() (returnValue string) {
	returnValue = ""
	counter := 0
	for counter <= 10 {
		returnValue = returnValue + strconv.FormatInt(int64(counter), 10)
		counter = counter + 1
	}
	return
}

// this is a forever loop, it is usful for waiting on events,
// have commented this out as it will break things
// func ForeverAlone(){
// 	for{
//
// 	}
// }

func IfStatementsAndTerntaryExpression(arg string) (return_item bool) {
	if arg == "yes" {
		return_item = true
	} else {
		return_item = false
	}
	return
}

// "short" statement, ie a statement can be run before the if
// usually this is done with error checking
func checkError() (*string, error) {
	if executeMethodResult := errors.New("hi"); executeMethodResult != nil {
		return nil, executeMethodResult
	}
	returnItem := "hi"
	return &returnItem, nil
}
