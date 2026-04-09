package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func NewNode(element int) *Node {
	return &Node{
		Val:  element,
		Next: nil,
	}
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Print() {
	curr := ll.Head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func main() {
	ll := NewLinkedList()

	ll.InsertRead(10)
	ll.InsertRead(20)
	ll.InsertRead(30)
	ll.InsertRead(40)
	ll.InsertRead(50)
	ll.InsertRead(60)
	ll.InsertRead(70)
	ll.InsertRead(80)
	ll.InsertRead(90)

	ll.InsertFront(0)

	ll.DeleteNode(90)

	ll.Print()

	idx := ll.SearchNode(40)
	fmt.Println("The element 40 is in index position : ", idx)
}
