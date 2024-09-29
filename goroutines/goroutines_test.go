package goroutines

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRunningOrder(t *testing.T) {
	result := GoRoutineExample()
	assert.Equal(t, 36, result)
}

func Test_populateChannel(t *testing.T) {
	myChannel := make(chan int)
	go func() { populateChannel(myChannel) }()
	runningValue := 0
	for recieved := range myChannel {
		runningValue = runningValue + recieved
	}
	assert.Equal(t, 4950, runningValue)
}

func Test_InterestingSelect(t *testing.T) {
	result := InterestingSelect()
	assert.Equal(t, 45, result)
}
