package ds

import "sync"

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
