package ds

import (
	"fmt"
	"strings"
	"sync"
)

type stackItem interface {
}

type Stack struct {
	Items []stackItem
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
func (q *Stack) Push(item stackItem) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	defer q.incrementSize()
	q.Items = append(q.Items, item)
}

// Pop removes and returns the object
// at the top of the stack
func (q *Stack) Pop() stackItem {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.Size <= 0 {
		return nil
	}

	defer q.decrementSize()
	var item stackItem
	item, q.Items = q.Items[q.Size-1], q.Items[:q.Size-1]

	return item
}

// Peek returns the stackItem at the top
// of the stack without removing it
func (q *Stack) Peek() stackItem {
	if q.Size <= 0 {
		return nil
	}
	return q.Items[q.Size-1]
}

// ToString returns a string representation
// of the stack
func (q *Stack) ToString() string {
	var sb strings.Builder
	for i := 0; i < q.Size; i++ {
		sb.WriteString(fmt.Sprint(q.Items[i]))
	}
	return sb.String()
}
