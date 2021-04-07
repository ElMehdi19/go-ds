package ds

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type any interface{}

func seedList(list *SinglyLinkedList, elements []any) {
	for _, element := range elements {
		list.Append(&Node{Value: element})
	}
}

func TestSLLAppend(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}

	elements := []any{1, 9, 9, 8}
	seedList(&list, elements)

	assert.Equal(len(elements), list.Size)

	items := list.Items()

	for i := 0; i < len(items); i++ {
		assert.Equal(elements[i], items[i])
	}
}

func TestSLLAppendAsync(t *testing.T) {
	assert := assert.New(t)
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

	assert.Equal(len(elements), list.Size)

	items := list.Items()
	for i := 0; i < len(items); i++ {
		assert.Equal(len(elements), list.Size)
	}
}

func TestSLLDelete(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	elements := []any{1, 2, 3, 4}

	seedList(&list, elements)
	assert.NotNil(list.Delete(5), "want out of range error; got nil")

	testFn := func(list *SinglyLinkedList, elements []any) {
		assert.Equal(len(elements), list.Size)

		items := list.Items()
		for i := 0; i < list.Size; i++ {
			assert.Equal(elements[i], items[i])
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
	assert := assert.New(t)
	list := SinglyLinkedList{}
	elements := []any{1, 3, 3, 4}
	seedList(&list, elements)

	list.Remove(3)
	elements = []any{1, 4}
	assert.Equal(len(elements), list.Size)

	items := list.Items()
	for i := 0; i < list.Size; i++ {
		assert.Equal(elements[i], items[i])
	}

	list = SinglyLinkedList{}
	elements = []any{3, 3, 3, 4}
	for _, elem := range elements {
		list.Append(&Node{Value: elem})
	}

	list.Remove(3)
	assert.Equal(1, list.Size)
	assert.Equal(4, list.Head.Value)
}

func TestSSLPrepend(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}

	testFn := func(list *SinglyLinkedList, size, value int) {
		assert.Equal(size, list.Size)

		assert.Equal(value, list.Head.Value)
	}

	list.Prepend(&Node{Value: 1})
	testFn(&list, 1, 1)

	list.Prepend(&Node{Value: 2})
	testFn(&list, 2, 2)
}

func TestSSLReverse(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	elements := []any{1, 2, 3, 4}
	seedList(&list, elements)

	list.Reverse()
	assert.Equal(4, list.Size)

	items := list.Items()
	elements = []any{4, 3, 2, 1}
	for i := 0; i < list.Size; i++ {
		assert.Equal(elements[i], items[i])
	}
}

func TestSSLClear(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	seedList(&list, []any{1, 2, 3, 4})
	list.Clear()

	assert.Equal(0, list.Size)
	assert.Nil(list.Head, "list.Head still points to a node")
}

func TestSLLUnique(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	elements := []any{3, 3, 1, 8, 0, 3, 8}
	seedList(&list, elements)
	list.Unique()

	elements = []any{3, 1, 8, 0}
	assert.Equal(len(elements), list.Size)

	items := list.Items()
	for i := 0; i < list.Size; i++ {
		assert.Equal(elements[i], items[i])
	}
}

func TestSLLSwap(t *testing.T) {
	assert := assert.New(t)
	list := SinglyLinkedList{}
	elements := []any{1, 2, 3, 4}
	seedList(&list, elements)
	list.Swap(1, 3)

	assert.Equal(4, list.Size)

	items := list.Items()
	elements = []any{1, 4, 3, 2}

	for i := 0; i < list.Size; i++ {
		assert.Equal(elements[i], items[i])
	}
}
