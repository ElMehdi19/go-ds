package ds

import (
	"errors"
	"sync"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type SinglyLinkedList struct {
	Head  *Node
	Size  int
	mutex sync.Mutex
}

func (s *SinglyLinkedList) incrementSize() {
	s.Size += 1
}

func (s *SinglyLinkedList) decrementSize() {
	s.Size -= 1
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
	defer s.incrementSize()
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

func (s *SinglyLinkedList) Delete(index int) error {
	if index >= s.Size {
		return errors.New("index out of range")
	}

	defer s.decrementSize()
	if index == 0 {
		s.Head = s.Head.Next
		return nil
	}

	currentNode := s.Head
	var previousNode *Node
	for i := 0; i < index; i++ {
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	previousNode.Next = currentNode.Next

	return nil
}

func (s *SinglyLinkedList) Remove(value interface{}) {
	if s.isEmpty() {
		return
	}

	for {

		if s.Head.Value == value {
			s.Head = s.Head.Next
			s.decrementSize()
		} else {
			break
		}
	}

	var previousNode *Node
	currentNode := s.Head

	for currentNode != nil {
		if currentNode.Value == value {
			previousNode.Next = currentNode.Next
			currentNode = currentNode.Next
			s.decrementSize()
		} else {
			previousNode = currentNode
			currentNode = currentNode.Next
		}
	}
}
