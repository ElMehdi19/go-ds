package ds

import "sync"

type item interface{}

type Queue struct {
	Items []item
	Size  int
	mutex sync.Mutex
}

func (q *Queue) incrementSize() {
	q.Size++
}

func (q *Queue) Push(item item) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	defer q.incrementSize()
	q.Items = append(q.Items, item)
}
