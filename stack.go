package ds

import (
	"fmt"
	"strings"
	"sync"
)

type Stack struct {
	Head  *Node
	Size  int
	mutex sync.Mutex
}

func (q *Stack) incrementSize() {
	q.Size++
}

func (q *Stack) decrementSize() {
	q.Size--
}

// Push takes a stackItem parameter
// and inserts it at the top of the stack
func (q *Stack) Push(item Any) {
	q.mutex.Lock()
	defer func() {
		q.incrementSize()
		q.mutex.Unlock()
	}()
	q.Head = &Node{Value: item, Next: q.Head}
}

// Pop removes and returns the object
// at the top of the stack
func (q *Stack) Pop() Any {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.Size <= 0 {
		return nil
	}

	defer q.decrementSize()
	var tmp *Node
	tmp, q.Head = q.Head, q.Head.Next

	return tmp.Value
}

// Peek returns the stackItem at the top
// of the stack without removing it
func (q *Stack) Peek() Any {
	if q.Size <= 0 {
		return nil
	}
	return q.Head.Value
}

// ToString returns a string representation
// of the stack
func (q *Stack) ToString() string {
	var sb strings.Builder

	for node := q.Head; node != nil; node = node.Next {
		sb.WriteString(fmt.Sprint(node.Value))
	}
	return sb.String()
}
