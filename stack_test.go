package ds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	assert := assert.New(t)
	q := Stack{}
	assert.Equal(0, q.Size, "queue should be empty at init time")

	items := []int{1, 9, 9, 8}
	for i, item := range items {
		q.Push(item)
		t.Run(fmt.Sprintf("q.Push #%d", i), func(t *testing.T) {
			assert.Equal(item, q.Head.Value)
		})
	}
}

func TestStackPop(t *testing.T) {
	assert := assert.New(t)
	q := Stack{}

	items := []int{1, 9, 9, 8}
	for _, item := range items {
		q.Push(item)
	}

	items = []int{8, 9, 9, 1}
	for i, item := range items {
		t.Run(fmt.Sprintf("q.Pop #%d", i), func(t *testing.T) {
			assert.Equal(item, q.Pop(), "want %d; got %d")
		})
	}

	assert.Equal(0, q.Size, "q.Size should be 0 after removing all queue items")
	assert.Nil(q.Pop(), "q head should be nil after removing all queue items")
}

func TestStackPeek(t *testing.T) {
	assert := assert.New(t)
	q := Stack{}

	assert.Nil(q.Peek(), "q head should be nil at init time")

	testFn := func(item int) {
		q.Push(item)
		assert.Equal(item, q.Peek(), "q.Peek error: want %d; got %d", 1, q.Peek())
	}
	for _, item := range []int{1, 9, 9, 8} {
		testFn(item)
	}

	assert.Equal(4, q.Size, "want %d; got %d,", 4, q.Size)
}

func TestStackInverse(t *testing.T) {
	assert := assert.New(t)
	q := Stack{}

	testStr := "racecar"
	for _, c := range testStr {
		q.Push(string(c))
	}
	assert.Equal(testStr, q.ToString(), "want %s; got %s", testStr, q.ToString())
}
