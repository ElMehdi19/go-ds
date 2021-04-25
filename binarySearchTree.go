package ds

import (
	"math"
	"sync"
)

type BSTNode struct {
	Value  int
	Parent *BSTNode
	Left   *BSTNode
	Right  *BSTNode
}

type BST struct {
	Root  *BSTNode
	mutex sync.Mutex
}

func (t *BST) Insert(n *BSTNode) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.Root == nil {
		t.Root = n
		return
	}
	t.insert(t.Root, n)
}

func (t *BST) insert(parent, node *BSTNode) {
	if parent == nil {
		return
	}

	if parent.Value >= node.Value {
		if parent.Left == nil {
			parent.Left = node
			return
		}
		t.insert(parent.Left, node)
	} else {
		if parent.Right == nil {
			parent.Right = node
			return
		}
		t.insert(parent.Right, node)
	}
}

func (t *BST) InOrderTraversal(n *BSTNode) []int {
	if n == nil {
		return nil
	}

	values := []int{}
	values = append(values, t.InOrderTraversal(n.Left)...)
	values = append(values, n.Value)
	values = append(values, t.InOrderTraversal(n.Right)...)

	return values
}

func (t *BST) PreOrderTraversal(n *BSTNode) []int {
	if n == nil {
		return nil
	}

	values := []int{n.Value}
	values = append(values, t.PreOrderTraversal(n.Left)...)
	values = append(values, t.PreOrderTraversal(n.Right)...)
	return values
}

func (t *BST) PostOrderTraversal(n *BSTNode) []int {
	if n == nil {
		return nil
	}

	values := []int{}
	values = append(values, t.PostOrderTraversal(n.Left)...)
	values = append(values, t.PostOrderTraversal(n.Right)...)
	values = append(values, n.Value)
	return values
}

func (t *BST) BFS(n *BSTNode) []int {
	if n == nil {
		return nil
	}

	var nodes []int
	var queue []*BSTNode

	pop := func() *BSTNode {
		top := queue[0]
		queue = queue[1:]
		return top
	}

	queue = append(queue, n)

	for len(queue) != 0 {
		qTop := pop()
		nodes = append(nodes, qTop.Value)

		if qTop.Left != nil {
			queue = append(queue, qTop.Left)
		}

		if qTop.Right != nil {
			queue = append(queue, qTop.Right)
		}
	}

	return nodes
}

func (t *BST) Height(n *BSTNode) int {
	if n == nil {
		return 0
	}

	currentHeight := 1 + math.Max(float64(t.Height(n.Left)), float64(t.Height(n.Right)))
	return int(currentHeight)
}

func (t *BST) Clear() {
	if t.Root == nil {
		return
	}
	t.Root = nil
}
