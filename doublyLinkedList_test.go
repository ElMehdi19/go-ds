package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDLLAppend(t *testing.T) {
	assert := assert.New(t)
	var list DoublyLinkedList

	assert.Nil(list.Head)

	elements := []Any{1, 9, 9, 8}
	seedList(&list, elements)

	assert.Equal(len(elements), list.Size())

	items := list.Items()
	for i := 0; i < list.Size(); i++ {
		assert.Equal(elements[i], items[i])
	}
}

func TestDLLGet(t *testing.T) {
	assert := assert.New(t)
	var list DoublyLinkedList

	item, err := list.Get(0)
	assert.Nil(item)
	assert.Error(err)

	elements := []Any{1, 9, 9, 8}
	seedList(&list, elements)

	for i := 0; i < len(elements); i++ {
		item, err = list.Get(i)
		assert.Nil(err)
		assert.Equal(elements[i], item)
	}

	item, err = list.Get(len(elements))
	assert.Nil(item)
	assert.Error(err)
}

func TestDLLPrepend(t *testing.T) {
	assert := assert.New(t)
	var list DoublyLinkedList

	list.Prepend(&Node{Value: 1})
	item, _ := list.Get(0)
	assert.Equal(1, item)

	list.Prepend(&Node{Value: 0})
	item, _ = list.Get(0)
	assert.Equal(0, item)

	item, _ = list.Get(1)
	assert.Equal(1, item)
}

func TestDLLDelete(t *testing.T) {
	assert := assert.New(t)
	var list DoublyLinkedList

	assert.Error(list.Delete(0))

	list.Append(&Node{Value: 1})
	assert.Equal(1, list.Size())
	assert.Error(list.Delete(1))
	assert.Nil(list.Delete(0))
	assert.Equal(0, list.Size())

	elements := []Any{1, 9, 9, 8}
	seedList(&list, elements)

	assert.Nil(list.Delete(1))
	assert.Nil(list.Delete(2))
	assert.Equal(2, list.Size())

	elements = elements[:2]

	for i := 0; i < len(elements); i++ {
		item, _ := list.Get(i)
		assert.Equal(elements[i], item)
	}
}

func TestDLLRemove(t *testing.T) {
	assert := assert.New(t)
	var list DoublyLinkedList

	elements := []Any{1, 9, 9, 8}
	seedList(&list, elements)
	assert.Equal(4, list.Size())

	list.Remove(9)
	assert.Equal(2, list.Size())

	item, _ := list.Get(0)
	assert.Equal(1, item)

	item, _ = list.Get(1)
	assert.Equal(8, item)

	list.Remove(1)
	assert.Equal(1, list.Size())
	item, _ = list.Get(0)
	assert.Equal(8, item)

	list.Remove(8)
	assert.Equal(0, list.Size())
}

func TestDLLSwap(t *testing.T) {
	assert := assert.New(t)
	var list DoublyLinkedList
	assert.Error(list.Swap(1, 2))

	elements := []Any{1, 2, 3, 4}
	seedList(&list, elements)

	items := list.Items()
	assert.Equal(len(elements), len(items))

	for i := 0; i < len(elements); i++ {
		assert.Equal(elements[i], items[i])
	}

	list.Swap(1, 2)
	elements = []Any{1, 3, 2, 4}
	items = list.Items()

	for i := 0; i < len(elements); i++ {
		assert.Equal(elements[i], items[i])
	}

	list.Swap(0, 3)
	elements = []Any{4, 3, 2, 1}
	items = list.Items()
	assert.Equal(len(elements), len(items))

	for i := 0; i < len(elements); i++ {
		assert.Equal(elements[i], items[i])
	}
}
