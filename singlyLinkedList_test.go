package ds

import (
	"sync"
	"testing"
)

func TestSLLAppend(t *testing.T) {
	list := SinglyLinkedList{}
	elements := []int{1, 9, 9, 8}
	for _, element := range elements {
		list.Append(&Node{Value: element})
	}

	if list.Size != len(elements) {
		t.Errorf("SinglyLinkedList.Size error: want %d; got %d", len(elements), list.Size)
	}

	items := list.Items()

	for i := 0; i < len(items); i++ {
		if items[i] != elements[i] {
			t.Errorf("node %d: want %d; got %d", i, elements[i], items[i])
		}
	}
}

func TestSLLAppendAsync(t *testing.T) {
	list := SinglyLinkedList{}
	elements := []int{1, 2, 3, 4}

	wg := &sync.WaitGroup{}
	wg.Add(len(elements))

	for _, element := range elements {
		go func(element int, wg *sync.WaitGroup) {
			list.Append(&Node{Value: element})
			wg.Done()
		}(element, wg)
	}

	wg.Wait()

	if list.Size != len(elements) {
		t.Errorf("SinglyLinkedList.Size error: want %d; got %d", len(elements), list.Size)
	}

	items := list.Items()
	for i := 0; i < len(items); i++ {
		if items[i] != elements[i] {
			t.Errorf("node %d: want %d; got %d", i, elements[i], items[i])
		}
	}
}
