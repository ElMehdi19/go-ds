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
