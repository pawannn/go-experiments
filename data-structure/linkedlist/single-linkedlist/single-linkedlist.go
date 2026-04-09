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

func (ll *LinkedList) AddNode(element int) {
	if ll.Head == nil {
		ll.Head = NewNode(element)
	} else {
		curr := ll.Head
		for curr.Next != nil {
			curr = curr.Next
		}

		curr.Next = NewNode(element)
	}
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

	ll.AddNode(10)
	ll.AddNode(20)
	ll.AddNode(30)
	ll.AddNode(40)
	ll.AddNode(50)
	ll.AddNode(60)
	ll.AddNode(70)
	ll.AddNode(80)
	ll.AddNode(90)

	ll.Print()
}
