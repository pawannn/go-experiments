package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func NewNode(element int) *Node {
	return &Node{
		Val:  element,
		Next: nil,
		Prev: nil,
	}
}

type CircularLinkedList struct {
	Head *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{}
}

func (cl *CircularLinkedList) Print() {
	curr := cl.Head

	for curr != nil {
		fmt.Printf("%d ", curr.Val)

		curr = curr.Next
		if curr == cl.Head {
			break
		}

		fmt.Printf(" -> ")
	}
}

func main() {
	cl := NewCircularLinkedList()

	cl.InsertFront(10)
	cl.InsertFront(20)
	cl.InsertFront(30)

	cl.InsertRear(40)

	cl.DeleteNode(30)

	exist := cl.NodeExist(10)
	fmt.Println("The element 10 : ", exist)

	cl.Print()
}
