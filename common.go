package ds

type Node struct {
	Value    interface{}
	Next     *Node
	Previous *Node
}

type Any interface{}

type List interface {
	Append(*Node)
	Get(int) (Any, error)
}

func seedList(list List, elements []Any) {
	for _, element := range elements {
		list.Append(&Node{Value: element})
	}
}
