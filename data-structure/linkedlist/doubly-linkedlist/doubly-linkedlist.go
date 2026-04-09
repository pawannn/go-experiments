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

type DoublyLinkedList struct {
	Head *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
	}
}

func (dl *DoublyLinkedList) PrintForward() {
	curr := dl.Head

	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
		fmt.Printf("->")
	}

	fmt.Println(nil)
}

func (dl *DoublyLinkedList) PrintBackward() {
	curr := dl.Head

	for curr.Next != nil {
		curr = curr.Next
	}

	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Prev
		fmt.Printf("<-")
	}

	fmt.Println(nil)
}

func main() {
	dl := NewDoublyLinkedList()

	dl.InsertFront(10)
	dl.InsertFront(20)
	dl.InsertFront(30)
	dl.InsertFront(40)
	dl.InsertFront(50)
	dl.InsertFront(60)
	dl.InsertFront(70)
	dl.InsertFront(80)

	dl.InsertRear(0)

	dl.DeleteNode(0)

	dl.PrintForward()
	dl.PrintBackward()

	idx := dl.SearchNode(40)
	fmt.Println("element 40 is in position : ", idx)
}
