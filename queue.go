package ds

import "sync"

type Queue struct {
	Top   *Node
	size  int
	mutex sync.Mutex
}

func (q *Queue) Size() int {
	size := 0
	if q.Top == nil {
		return size
	}

	for node := q.Top; node != nil; node = node.Next {
		size += 1
	}

	return size
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) incrementSize() {
	q.size++
}

func (q *Queue) decrementSize() {
	if q.size == 0 {
		return
	}
	q.size--
}

func (q Queue) Peek() Any {
	if q.IsEmpty() {
		return nil
	}
	return q.Top.Value
}

func (q *Queue) Push(item Any) {
	if item == nil {
		return
	}
	q.mutex.Lock()
	defer func() {
		q.incrementSize()
		q.mutex.Unlock()
	}()

	if q.IsEmpty() {
		q.Top = &Node{Value: item}
		return
	}

	node := q.Top
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &Node{Value: item}
}
