package gameobject

// func TestLinkedList(t *testing.T) {
// 	list := InitLinkedList()

// 	list.Append(mapobjects.Wall{})
// 	list.Append(mapobjects.Wall{})
// 	list.Append(mapobjects.Wall{})
// 	list.Append(mapobjects.Wall{})
// 	a := mapobjects.InitWall(10, "TOP", rl.Rectangle{})
// 	head := mapobjects.InitWall(10, "TOP", rl.Rectangle{})
// 	list.Append(a)
// 	list.Append(head)

// 	if list.Head.Value.GetPosition() != head.Pos {
// 		t.Error("Head different")
// 	}

// 	if list.Head.Next.Value.GetPosition() != a.Pos {
// 		t.Error("Next different")
// 	}

// }

// func TestLinkedListDelete(t *testing.T) {
// 	list := InitLinkedList()

// 	list.Append(mapobjects.Wall{})
// 	list.Append(mapobjects.Wall{})
// 	list.Append(mapobjects.Wall{})
// 	list.Append(mapobjects.Wall{})
// 	a := mapobjects.Wall{}
// 	head := mapobjects.InitWall(10, "TOP", rl.Rectangle{})
// 	list.Append(a)
// 	list.Append(head)

// 	list.Head.Next.Unlink(list)

// 	if list.Head.Next.Value.GetPosition() == a.Pos {
// 		t.Error("Not deleted")
// 	}

// 	list.Head.Unlink(list)

// 	if list.Head.Value.GetPosition() == head.Pos {
// 		t.Error("Not deleted head")
// 	}

// 	fmt.Print(list.Head.Value)

// }
