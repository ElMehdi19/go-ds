package ds

type Node struct {
	Value    interface{}
	Next     *Node
	Previous *Node
}

type Any interface{}

type List interface {
	Append(n *Node)
}

func seedList(list List, elements []Any) {
	for _, element := range elements {
		list.Append(&Node{Value: element})
	}
}
