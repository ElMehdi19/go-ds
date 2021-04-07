package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func seedQueue(q *Queue, items ...Any) {
	for _, item := range items {
		q.Push(item)
	}
}

func TestQueuePush(t *testing.T) {
	assert := assert.New(t)

	queue := Queue{}
	assert.Nil(queue.Peek())

	// Pushing to empty queue
	queue.Push(1)
	queue.Push(2)
	// Pushing nil item
	queue.Push(nil)
	assert.Equal(1, queue.Peek())
	assert.Equal(2, queue.Size())

	// Pushing to non empty queue
	queue = Queue{Top: &Node{Value: 3}}
	assert.Equal(3, queue.Peek())
	assert.Equal(1, queue.Size())
}

func TestQueuePop(t *testing.T) {
	assert := assert.New(t)

	queue := Queue{}
	elements := []Any{1, 2, 3}

	seedQueue(&queue, elements...)
	for _, item := range elements {
		assert.Equal(item, queue.Pop())
	}

	assert.Zero(queue.Size())
	assert.Nil(queue.Pop())
}
