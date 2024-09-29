package goroutines

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"golang.org/x/tour/tree"
)

func Test_Walk(t *testing.T) {
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
	}()
	runningValue := 0
	for sentFromTree := range ch {
		runningValue = runningValue + sentFromTree
	}
}

func Test_Same(t *testing.T) {
	treeOne := tree.New(1)
	treeTwo := tree.New(1)
	assert.Equal(t, Same(treeOne, treeTwo), false)
	assert.Equal(t, Same(treeOne, treeOne), true)
}

func Test_Syncrhonized(t *testing.T) {
	unit := Mine{}
	doneOne := make(chan int)
	doneTwo := make(chan int)
	go func() {
		unit.Add(12)
		doneOne <- 0
	}()
	go func() {
		unit.Add(42)
		doneTwo <- 0
	}()
	<-doneOne
	<-doneTwo
	assert.Equal(t, 54, unit.value)
}
