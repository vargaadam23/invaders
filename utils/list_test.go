package utils

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := InitLinkedList()

	a := 3
	b := 4
	list.Append(a)
	list.Append(b)
	list.Append(4)
	list.Append(5)
	list.Append("aaa")
	list.Append("vvv")

	fmt.Printf("head=%v", *list.Head.Value)
	fmt.Printf("next=%v", *list.Head.Next.Next.Value)
	fmt.Printf("next=%v", *list.Head.Next.Next.Next.Value)
	fmt.Printf("next=%v", *list.Head.Next.Next.Next.Next.Value)

}
