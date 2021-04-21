package ds

import (
	"errors"
	"fmt"
	"sync"
)

type SinglyLinkedList struct {
	Head  *Node
	size  int
	mutex sync.Mutex
}

func (s *SinglyLinkedList) incrementSize() {
	s.size += 1
}

func (s *SinglyLinkedList) decrementSize() {
	s.size -= 1
}

func (s *SinglyLinkedList) Size() int {
	return s.size
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

// Append takes a *Node and and adds it to the list
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

// Delete takes a 0-based index and deletes
// its correspendant node from the list
// if it does exist otherwise returns a non-nil error
func (s *SinglyLinkedList) Delete(index int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if index >= s.size {
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

// Remove takes a value and deletes all the nodes
// with the same value in the list
func (s *SinglyLinkedList) Remove(value Any) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
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

// Prepend takes a *Node and inserts it
// at the head of the list
func (s *SinglyLinkedList) Prepend(n *Node) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	defer s.incrementSize()
	if s.isEmpty() {
		s.Head = n
		return
	}

	n.Next = s.Head
	s.Head = n
}

// Reverse reverses the nodes order in the list
func (s *SinglyLinkedList) Reverse() {
	if s.size <= 1 {
		return
	}

	currentNode := s.Head
	var previousNode *Node
	var tempNode *Node

	for currentNode != nil {
		tempNode = currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = tempNode
	}
	s.Head = previousNode
}

// Clear removes all the list elements
func (s *SinglyLinkedList) Clear() {
	if s.isEmpty() {
		return
	}
	s.Head = nil
	s.size = 0
}

// Unique removes all the duplicated nodes
// based on their values
func (s *SinglyLinkedList) Unique() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.size <= 1 {
		return
	}
	visited := make(map[interface{}]bool)
	currentNode := s.Head
	var previousNode *Node

	for currentNode != nil {
		if _, ok := visited[currentNode.Value]; ok {
			previousNode.Next = currentNode.Next
			s.decrementSize()
		} else {
			visited[currentNode.Value] = true
			previousNode = currentNode
		}
		currentNode = previousNode.Next
	}
}

// Swap takes two int params and swap the position
// of their correspending nodes in the list
func (s *SinglyLinkedList) Swap(i, j int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if i >= s.size || j >= s.size {
		return fmt.Errorf("index out of range")
	}
	if i == j || s.size <= 1 {
		return nil
	}

	if i > j {
		i, j = j, i
	}

	currentNodeX := s.Head
	var previousNodeX *Node

	for x := 0; x < i; x++ {
		previousNodeX = currentNodeX
		currentNodeX = currentNodeX.Next
	}

	currentNodeY := s.Head
	var previousNodeY *Node

	for x := 0; x < j; x++ {
		previousNodeY = currentNodeY
		currentNodeY = currentNodeY.Next
	}

	if previousNodeX != nil {
		previousNodeX.Next = currentNodeY
	} else {
		s.Head = currentNodeY
	}

	if previousNodeY != nil {
		previousNodeY.Next = currentNodeX
	}

	currentNodeX.Next, currentNodeY.Next = currentNodeY.Next, currentNodeX.Next

	return nil

}

func (s *SinglyLinkedList) Get(id int) (Any, error) {
	if s.size == 0 || id >= s.size {
		return nil, fmt.Errorf("index out of range")
	}

	currentNode := s.Head

	for i := 0; i < id; i++ {
		currentNode = currentNode.Next
	}

	return currentNode.Value, nil
}
