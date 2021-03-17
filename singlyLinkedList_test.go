package ds

import "testing"

func TestSLLAppend(t *testing.T) {
	list := SinglyLinkedList{}
	elements := []int{1, 9, 9, 8}
	for _, element := range elements {
		list.Append(&Node{Value: element})
	}

	items := list.Items()
	if len(items) != len(elements) {
		t.Errorf("SinglyLinkedList.Items length error: want %d; got %d", len(items), len(elements))
	}

	for i := 0; i < len(items); i++ {
		if items[i] != elements[i] {
			t.Errorf("node %d: want %d; got %d", i, items[i], elements[i])
		}
	}
}
