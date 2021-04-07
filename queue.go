package ds

import (
	"fmt"
	"strings"
	"sync"
)

type Queue struct {
	// fields are only accessible by
	// readers: Peek, Pop
	top   *Node
	mutex sync.Mutex
}

func (q *Queue) Size() int {
	size := 0
	if q.top == nil {
		return size
	}

	for node := q.top; node != nil; node = node.Next {
		size += 1
	}

	return size
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Peek() Any {
	if q.IsEmpty() {
		return nil
	}
	return q.top.Value
}

func (q *Queue) Push(item Any) {
	if item == nil {
		return
	}
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		q.top = &Node{Value: item}
		return
	}

	node := q.top
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &Node{Value: item}
}

func (q *Queue) Pop() Any {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		return nil
	}

	top := q.top.Value
	q.top = q.top.Next
	return top
}

func (q *Queue) Clear() {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		return
	}

	q.top = nil
}

func (q *Queue) ToString() string {
	if q.IsEmpty() {
		return ""
	}

	var sb strings.Builder
	for node := q.top; node != nil; node = node.Next {
		sb.WriteString(fmt.Sprint(node.Value))
	}
	return sb.String()
}
