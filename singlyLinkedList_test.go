package ds

import (
	"sync"
	"testing"
)

type any interface{}

func seedList(list *SinglyLinkedList, elements []any) {
	for _, element := range elements {
		list.Append(&Node{Value: element})
	}
}

func TestSLLAppend(t *testing.T) {
	list := SinglyLinkedList{}
	elements := []any{1, 9, 9, 8}
	seedList(&list, elements)

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
	elements := []any{1, 2, 3, 4}
	seedList(&list, elements)

	if err := list.Delete(5); err == nil {
		t.Error("want out of range error; got nil")
	}

	testFn := func(list *SinglyLinkedList, elements []any) {
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
	elements = []any{2, 4}
	testFn(&list, elements)
}

func TestSLLRemove(t *testing.T) {
	list := SinglyLinkedList{}
	elements := []any{1, 3, 3, 4}
	seedList(&list, elements)

	list.Remove(3)
	elements = []any{1, 4}
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
	elements = []any{3, 3, 3, 4}
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
	elements := []any{1, 2, 3, 4}
	seedList(&list, elements)

	list.Reverse()
	if list.Size != 4 {
		t.Fatalf("list.Size error: want %d; got %d", 4, list.Size)
	}

	items := list.Items()
	elements = []any{4, 3, 2, 1}
	for i := 0; i < list.Size; i++ {
		if items[i] != elements[i] {
			t.Errorf("item #%d: want %d; got %d", i, elements[i], items[i])
		}
	}
}

func TestSSLClear(t *testing.T) {
	list := SinglyLinkedList{}
	seedList(&list, []any{1, 2, 3, 4})
	list.Clear()

	if list.Size != 0 {
		t.Errorf("list.Size error: want %d; got %d", 0, list.Size)
	}

	if list.Head != nil {
		t.Errorf("list.Head still points to a node")
	}
}
