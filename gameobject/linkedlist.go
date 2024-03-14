package gameobject

type Node struct {
	Value GameObject
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

func (list *LinkedList) Append(value GameObject) {
	if list.Head.Value == nil {
		list.Head.Value = value

		return
	}

	oldHead := list.Head
	list.Head = &Node{
		Next:  oldHead,
		Value: value,
	}

	oldHead.Prev = list.Head
}

func (node *Node) Unlink(list *LinkedList) *Node {
	if list.Head == node {
		list.Head = node.Next
	}

	next := node.Next

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	node.Next = nil
	node.Prev = nil
	node.Value = nil
	node = nil

	return next
}
