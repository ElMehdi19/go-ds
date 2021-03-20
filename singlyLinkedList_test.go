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

func TestSLLDelete(t *testing.T) {
	list := SinglyLinkedList{}
	elements := []int{1, 2, 3, 4}
	for _, elem := range elements {
		list.Append(&Node{Value: elem})
	}

	if err := list.Delete(5); err == nil {
		t.Error("want out of range error; got nil")
	}

	testFn := func(list *SinglyLinkedList, elements []int) {
		if list.Size != len(elements) {
			t.Fatalf("list.Size error: want %d; got %d", len(elements), list.Size)
		}

		items := list.Items()
		for i := 0; i < list.Size; i++ {
			if items[i] != elements[i] {
				t.Fatalf("item #%d: want %d; got %d", i, elements[i], items[i])
			}
		}
	}

	list.Delete(0)
	elements = elements[1:]
	testFn(&list, elements)

	list.Delete(1)
	elements = []int{2, 4}
	testFn(&list, elements)
}

func TestSLLRemove(t *testing.T) {
	list := SinglyLinkedList{}
	elements := []int{1, 3, 3, 4}
	for _, elem := range elements {
		list.Append(&Node{Value: elem})
	}

	list.Remove(3)
	elements = []int{1, 4}
	if list.Size != len(elements) {
		t.Fatalf("list.Size error: want %d; got %d", len(elements), list.Size)
	}

	items := list.Items()
	for i := 0; i < list.Size; i++ {
		if items[i] != elements[i] {
			t.Errorf("item #%d: want %d; got %d", i, elements[i], items[i])
		}
	}

	list = SinglyLinkedList{}
	elements = []int{3, 3, 3, 4}
	for _, elem := range elements {
		list.Append(&Node{Value: elem})
	}

	list.Remove(3)
	if list.Size != 1 {
		t.Errorf("list.Size error: want %d; got %d", 1, list.Size)
	}

	if list.Head.Value != 4 {
		t.Errorf("want %d; got %d", 4, list.Head.Value)
	}
}

func TestSSLPrepend(t *testing.T) {
	list := SinglyLinkedList{}

	testFn := func(list *SinglyLinkedList, size, value int) {
		if list.Size != size {
			t.Errorf("list.Size error: want %d; got %d", 1, list.Size)
		}

		if list.Head.Value != value {
			t.Errorf("want %d; got %d", 1, list.Head.Value)
		}
	}

	list.Prepend(&Node{Value: 1})
	testFn(&list, 1, 1)

	list.Prepend(&Node{Value: 2})
	testFn(&list, 2, 2)
}

func TestSSLReverse(t *testing.T) {
	list := SinglyLinkedList{}
	elements := []int{1, 2, 3, 4}
	for _, element := range elements {
		list.Append(&Node{Value: element})
	}

	list.Reverse()
	if list.Size != 4 {
		t.Fatalf("list.Size error: want %d; got %d", 4, list.Size)
	}

	items := list.Items()
	elements = []int{4, 3, 2, 1}
	for i := 0; i < list.Size; i++ {
		if items[i] != elements[i] {
			t.Errorf("item #%d: want %d; got %d", i, elements[i], items[i])
		}
	}
}
