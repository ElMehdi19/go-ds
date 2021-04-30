package ds

import (
	"fmt"
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

func TestBSTBreadFirstSearch(t *testing.T) {
	assert := assert.New(t)
	var tree BST

	elements := []int{5, 2, 8, 1, 3}
	seedBST(&tree, elements)

	assert.NotNil(tree.Root)
	items := tree.BFS(tree.Root)

	assert.Equal(len(elements), len(items))
	for i := 0; i < len(items); i++ {
		assert.Equal(elements[i], items[i])
	}
}

func TestBSTHeight(t *testing.T) {
	assert := assert.New(t)
	var tree BST

	assert.Zero(tree.Height(tree.Root))

	elements := []int{5, 2, 8, 1, 3}
	seedBST(&tree, elements)

	assert.Equal(3, tree.Height(tree.Root))

	elements = []int{4, 7, 9}
	seedBST(&tree, elements)

	assert.Equal(4, tree.Height(tree.Root))

	elements = []int{0, 2}
	seedBST(&tree, elements)
	assert.Equal(4, tree.Height(tree.Root))
}

func TestMinBSTFromArray(t *testing.T) {
	assert := assert.New(t)
	items := []int{2, 4, 1, 9, 5, 8}
	var tree BST

	root := tree.MinBSTFromArray(items)
	assert.NotNil(root)

	assert.Equal(3, tree.Height(root))

	inOrder := tree.InOrderTraversal(root)
	assert.Equal("[1 2 4 5 8 9]", fmt.Sprintf("%v", inOrder), inOrder)

	preOrder := tree.PreOrderTraversal(root)
	assert.Equal("[4 1 2 8 5 9]", fmt.Sprintf("%v", preOrder), preOrder)

	postOrder := tree.PostOrderTraversal(root)
	assert.Equal("[2 1 5 9 8 4]", fmt.Sprintf("%v", postOrder), postOrder)

	/* [1, 2, 4, 5, 8, 9]

		 4
	   /  \
	  1    8
	   \  / \
	    2 5 9

	*/

}
