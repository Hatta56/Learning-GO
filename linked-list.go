package main

import "fmt"

type Node struct {
	Prev *Node
	Next *Node
	Key  interface{}
}
type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		Next: L.head,
		Key:  key,
	}
	if L.head != nil {
		L.head.Prev = list
	}
	L.head = list
	l := L.head
	for l.Next != nil {
		l = l.Next
	}
	L.tail = l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%v ", list.Key)
		list = list.Next
	}
	fmt.Println()
}
func main() {
	link := List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(18)

	fmt.Printf("Head : %v\n", link.head.Key)
	fmt.Printf("Tail : %v\n", link.tail.Key)
	link.Display()
}
