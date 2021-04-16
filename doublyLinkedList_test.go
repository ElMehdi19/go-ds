package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDLLAppend(t *testing.T) {
	assert := assert.New(t)
	list := DoublyLinkedList{}

	elements := []Any{1, 9, 9, 8}
	seedList(&list, elements)

	assert.Equal(len(elements), list.Size())

	items := list.Items()
	for i := 0; i < list.Size(); i++ {
		assert.Equal(elements[i], items[i])
	}
}
