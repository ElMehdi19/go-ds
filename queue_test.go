package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
