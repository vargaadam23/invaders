package utils

type Node struct {
	Value *interface{}
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Head *Node
}

func InitLinkedList() *LinkedList {
	return &LinkedList{
		Head: &Node{
			Next:  nil,
			Prev:  nil,
			Value: nil,
		},
	}
}

func (list *LinkedList) Append(value interface{}) {
	if list.Head.Value == nil {
		list.Head.Value = &value

		return
	}

	oldHead := list.Head
	list.Head = &Node{
		Next:  oldHead,
		Value: &value,
	}

	oldHead.Prev = list.Head
}
