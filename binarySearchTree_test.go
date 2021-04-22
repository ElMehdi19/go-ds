package ds

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func seedBST(tree *BST, elements []int) {
	for _, element := range elements {
		tree.Insert(&BSTNode{Value: element})
	}
}

func TestBSTInsert(t *testing.T) {
	assert := assert.New(t)
	var tree BST

	elements := []int{5, 2, 8, 1, 3}
	seedBST(&tree, elements)

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

func TestBSTPreOrderTraversal(t *testing.T) {
	assert := assert.New(t)
	var tree BST

	elements := []int{5, 2, 8, 1, 3}
	seedBST(&tree, elements)

	assert.NotNil(tree.Root)
	elements = []int{5, 2, 1, 3, 8}
	items := tree.PreOrderTraversal(tree.Root)

	assert.Equal(len(elements), len(items))
	for i := 0; i < len(items); i++ {
		assert.Equal(elements[i], items[i])
	}
}

func TestBSTPostOrderTraversal(t *testing.T) {
	assert := assert.New(t)
	var tree BST

	elements := []int{5, 2, 8, 1, 3}
	seedBST(&tree, elements)

	assert.NotNil(tree.Root)
	elements = []int{1, 3, 2, 8, 5}
	items := tree.PostOrderTraversal(tree.Root)

	assert.Equal(len(elements), len(items))
	for i := 0; i < len(items); i++ {
		assert.Equal(elements[i], items[i])
	}
}
