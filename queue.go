package ds

import "sync"

type Queue struct {
	Top   *Node
	size  int
	mutex sync.Mutex
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
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
