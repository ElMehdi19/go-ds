package ds

import (
	"fmt"
	"strings"
	"sync"
)

type Stack struct {
	top   *Node
	mutex sync.Mutex
}

// Push takes a stackItem parameter
// and inserts it at the top of the stack
func (q *Stack) Push(item Any) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.top = &Node{Value: item, Next: q.top}
}

// Pop removes and returns the object
// at the top of the stack
func (q *Stack) Pop() Any {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		return nil
	}

	var tmp *Node
	tmp, q.top = q.top, q.top.Next

	return tmp.Value
}

// Peek returns the stackItem at the top
// of the stack without removing it
func (q *Stack) Peek() Any {
	if q.IsEmpty() {
		return nil
	}
	return q.top.Value
}

// ToString returns a string representation
// of the stack
func (q *Stack) ToString() string {
	var sb strings.Builder

	for node := q.top; node != nil; node = node.Next {
		sb.WriteString(fmt.Sprint(node.Value))
	}
	return sb.String()
}

func (q *Stack) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Stack) Size() int {
	size := 0

	if q.top == nil {
		return size
	}

	for node := q.top; node != nil; node = node.Next {
		size += 1
	}

	return size
}
