package character

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := InitLinkedList()

	list.Append(&Bullet{})
	list.Append(&Bullet{})
	list.Append(&Bullet{})
	list.Append(&Bullet{})
	a := &Bullet{}
	head := &Bullet{}
	list.Append(a)
	list.Append(head)

	if list.Head.Value != head {
		t.Error("Head different")
	}

	if list.Head.Next.Value != a {
		t.Error("Next different")
	}

}

func TestLinkedListDelete(t *testing.T) {
	list := InitLinkedList()

	list.Append(&Bullet{})
	list.Append(&Bullet{})
	list.Append(&Bullet{})
	list.Append(&Bullet{})
	a := &Bullet{}
	head := &Bullet{}
	list.Append(a)
	list.Append(head)

	list.Head.Next.Unlink(list)

	if list.Head.Next.Value == a {
		t.Error("Not deleted")
	}

	list.Head.Unlink(list)

	if list.Head.Value == head {
		t.Error("Not deleted head")
	}

	fmt.Print(*list.Head.Value)

}
