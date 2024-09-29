package goroutines

import (
	"sync"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	recurssiveWalk(t, ch)
	close(ch)
}
func recurssiveWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		ch <- t.Value
		if t.Left != nil {
			recurssiveWalk(t.Left, ch)
		}
		if t.Right != nil {
			recurssiveWalk(t.Right, ch)
		}
	}
}

func Same(t1, t2 *tree.Tree) bool {
	treeOneChannel, treeTwoChannel := make(chan int), make(chan int)
	go Walk(t1, treeOneChannel)
	go Walk(t2, treeTwoChannel)
	for {
		treeOneElement, treeOneChannelClosed := <-treeOneChannel
		treeTwoElement, treeTwoChannelClosed := <-treeTwoChannel
		if treeOneElement != treeTwoElement || treeOneChannelClosed != treeTwoChannelClosed {
			return false
		}
		if !treeOneChannelClosed {
			break
		}
	}
	return true
}

//Java Synchronised word is mutex

type Mine struct {
	lock  sync.Mutex
	value int
}

func (m *Mine) Add(arg int) {
	m.lock.Lock()
	m.value = m.value + arg
	m.lock.Unlock()
}
