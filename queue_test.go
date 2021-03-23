package ds

import (
	"fmt"
	"testing"
)

func TestQPush(t *testing.T) {
	q := Queue{}
	if q.Size != 0 {
		t.Errorf("queue should be empty at init time")
	}

	items := []int{2, 3, 1, 9}
	for i, item := range items {
		q.Push(item)
		t.Run(fmt.Sprintf("q.Push #%d", i), func(t *testing.T) {
			if item != q.Items[q.Size-1] {
				t.Errorf("want %d; got %d", item, q.Items[q.Size-1])
			}
		})
	}
}

func TestQPop(t *testing.T) {
	q := Queue{}

	items := []int{2, 3, 1, 9}
	for _, item := range items {
		q.Push(item)
	}

	items = []int{9, 1, 3, 2}
	for i, item := range items {
		t.Run(fmt.Sprintf("q.Pop #%d", i), func(t *testing.T) {
			top := q.Pop()
			if item != top {
				t.Errorf("want %d; got %d", item, top)
			}
		})
	}

	if q.Size != 0 {
		t.Errorf("q.Size should be 0 after removing all queue items")
	}

	if q.Pop() != nil {
		t.Errorf("q.Pop should return nil after removing all queue items")
	}
}
