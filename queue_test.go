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
