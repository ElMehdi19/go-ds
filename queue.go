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

func (q *Queue) decrementSize() {
	q.Size--
}

func (q *Queue) Push(item item) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	defer q.incrementSize()
	q.Items = append(q.Items, item)
}

func (q *Queue) Pop() item {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.Size <= 0 {
		return nil
	}

	defer q.decrementSize()
	var item item
	item, q.Items = q.Items[q.Size-1], q.Items[:q.Size-1]

	return item
}

func (q *Queue) Peek() item {
	if q.Size <= 0 {
		return nil
	}
	return q.Items[q.Size-1]
}
