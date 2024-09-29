package typesstructcsslicesmaps

import (
	"strings"
)

func WordCount(s string) map[string]int {
	returningValue := map[string]int{}
	for _, value := range strings.Fields(s) {
		if foundValue, isPresent := returningValue[value]; isPresent {
			returningValue[value] = foundValue + 1
		} else {
			returningValue[value] = 1
		}
	}
	return returningValue
}

func fibonacci() func() int {
	lastTwoNumbers := [2]int{}
	return func() int {
		if lastTwoNumbers[0] == 0 && lastTwoNumbers[1] == 0 {
			lastTwoNumbers[1] = 1
			return 0
		} else {
			newValue := lastTwoNumbers[0] + lastTwoNumbers[1]
			lastTwoNumbers[0] = lastTwoNumbers[1]
			lastTwoNumbers[1] = newValue
			return lastTwoNumbers[0]
		}
	}
}

type MyReader struct{}

func (m MyReader) Read(bytes []byte) (int, error) {
	var bytesLength = len(bytes)
	for counter := 0; counter < bytesLength; counter++ {
		bytes[counter] = 'A'
	}
	return bytesLength, nil
}
