package ds

import "sync"

type DoublyLinkedList struct {
	Head  *Node
	size  int
	mutex sync.Mutex
}

func (d *DoublyLinkedList) incrementSize() {
	d.size += 1
}

func (d *DoublyLinkedList) Size() int {
	return d.size
}

func (d *DoublyLinkedList) Items() []Any {
	var items []Any

	for currentNode := d.Head; currentNode != nil; currentNode = currentNode.Next {
		items = append(items, currentNode.Value)
	}

	return items
}

func (d *DoublyLinkedList) Append(n *Node) {
	d.mutex.Lock()
	defer d.incrementSize()
	defer d.mutex.Unlock()

	if d.size == 0 {
		d.Head = n
		return
	}

	currentNode := d.Head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	n.Previous = currentNode
	currentNode.Next = n
}
