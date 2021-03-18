package ds

import "sync"

type Node struct {
	Value interface{}
	Next  *Node
}

type SinglyLinkedList struct {
	Head  *Node
	Size  int
	mutex sync.Mutex
}

func (s *SinglyLinkedList) addToSize(n int) {
	s.Size += n
}

func (s *SinglyLinkedList) isEmpty() bool {
	return s.Head == nil
}

func (s *SinglyLinkedList) Items() []interface{} {
	var items []interface{}

	for currentNode := s.Head; currentNode != nil; currentNode = currentNode.Next {
		items = append(items, currentNode.Value)
	}

	return items
}

func (s *SinglyLinkedList) Append(n *Node) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	defer s.addToSize(1)
	if s.isEmpty() {
		s.Head = n
		return
	}

	currentNode := s.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = n
}
