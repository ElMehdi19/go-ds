package ds

import (
	"fmt"
	"strings"
	"sync"
)

type Stack struct {
	Top   *Node
	size  int
	mutex sync.Mutex
}

func (q *Stack) incrementSize() {
	q.size++
}

func (q *Stack) decrementSize() {
	q.size--
}

// Push takes a stackItem parameter
// and inserts it at the top of the stack
func (q *Stack) Push(item Any) {
	q.mutex.Lock()
	defer func() {
		q.incrementSize()
		q.mutex.Unlock()
	}()
	q.Top = &Node{Value: item, Next: q.Top}
}

// Pop removes and returns the object
// at the top of the stack
func (q *Stack) Pop() Any {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		return nil
	}

	defer q.decrementSize()
	var tmp *Node
	tmp, q.Top = q.Top, q.Top.Next

	return tmp.Value
}

// Peek returns the stackItem at the top
// of the stack without removing it
func (q *Stack) Peek() Any {
	if q.IsEmpty() {
		return nil
	}
	return q.Top.Value
}

// ToString returns a string representation
// of the stack
func (q *Stack) ToString() string {
	var sb strings.Builder

	for node := q.Top; node != nil; node = node.Next {
		sb.WriteString(fmt.Sprint(node.Value))
	}
	return sb.String()
}

func (q *Stack) IsEmpty() bool {
	return q.size == 0
}

func (q *Stack) Size() int {
	return q.size
}
