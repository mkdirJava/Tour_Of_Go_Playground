package typesstructcsslicesmaps

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"golang.org/x/tour/reader"
)

func Test_WordCount(t *testing.T) {
	assert.Equal(t, map[string]int{"Hi": 1, "There": 2, "World": 1}, WordCount("Hi There There World"))
}

func Test_fibonacci(t *testing.T) {
	fibonacciFunc := fibonacci()
	result := []int{fibonacciFunc(), fibonacciFunc(), fibonacciFunc(), fibonacciFunc(), fibonacciFunc(), fibonacciFunc()}
	assert.Equal(t, []int{0, 1, 1, 2, 3, 5}, result)
}

func Test_myReader(t *testing.T) {
	myReader := MyReader{}
	reader.Validate(myReader)
	bytes := [10]byte{}
	readBytes, err := myReader.Read(bytes[:])
	assert.Equal(t, nil, err)
	assert.Equal(t, 10, readBytes)
}
