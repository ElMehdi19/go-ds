package ds

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTInsert(t *testing.T) {
	assert := assert.New(t)
	var tree BST

	elements := []int{5, 2, 8, 1, 3}
	for _, element := range elements {
		tree.Insert(&BSTNode{Value: element})
	}

	assert.NotNil(tree.Root)

	items := tree.InOrderTraversal(tree.Root)
	sort.Slice(elements, func(i, j int) bool {
		return elements[i] < elements[j]
	})

	assert.Equal(len(elements), len(items))

	for i := 0; i < len(items); i++ {
		assert.Equal(elements[i], items[i])
	}
}
