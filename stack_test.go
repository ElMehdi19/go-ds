package ds

import (
	"fmt"
	"testing"
)

func TestStackPush(t *testing.T) {
	q := Stack{}
	if q.Size != 0 {
		t.Errorf("queue should be empty at init time")
	}

	items := []int{1, 9, 9, 8}
	for i, item := range items {
		q.Push(item)
		t.Run(fmt.Sprintf("q.Push #%d", i), func(t *testing.T) {
			if item != q.Items[q.Size-1] {
				t.Errorf("want %d; got %d", item, q.Items[q.Size-1])
			}
		})
	}
}

func TestStackPop(t *testing.T) {
	q := Stack{}

	items := []int{1, 9, 9, 8}
	for _, item := range items {
		q.Push(item)
	}

	items = []int{8, 9, 9, 1}
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
		t.Errorf("q head should be nil after removing all queue items")
	}
}

func TestStackPeek(t *testing.T) {
	q := Stack{}
	if q.Peek() != nil {
		t.Errorf("q head should be nil at init time")
	}

	testFn := func(item int) {
		q.Push(item)
		if head := q.Peek(); head != item {
			t.Errorf("q.Peek error: want %d; got %d", 1, head)
		}
	}
	for _, item := range []int{1, 9, 9, 8} {
		testFn(item)
	}

	if q.Size != 4 {
		t.Errorf("want %d; got %d,", 4, q.Size)
	}
}

func TestStackInverse(t *testing.T) {
	q := Stack{}
	testStr := "racecar"
	for _, c := range testStr {
		q.Push(string(c))
	}
	if q.ToString() != testStr {
		t.Errorf("want %s; got %s", testStr, q.ToString())
	}
}
