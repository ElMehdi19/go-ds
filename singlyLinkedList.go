package ds

type Node struct {
	Value interface{}
	Next  *Node
}

type SinglyLinkedList struct {
	Head *Node
	size int
}

func (s *SinglyLinkedList) addToSize(n int) {
	s.size += n
}

func (s *SinglyLinkedList) Append(n *Node) {
	defer s.addToSize(1)
	if s.Head == nil {
		s.Head = n
		return
	}

	currentNode := s.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = n
}

func (s SinglyLinkedList) Items() []interface{} {
	var items []interface{}

	for currentNode := s.Head; currentNode != nil; currentNode = currentNode.Next {
		items = append(items, currentNode.Value)
	}

	return items
}
