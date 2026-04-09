package main

import "fmt"

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

func NewNode(element int) *Node {
	return &Node{
		Val:   element,
		Right: nil,
		Left:  nil,
	}
}

func InsertNode(node *Node, element int) *Node {
	if node == nil {
		return NewNode(element)
	}

	if element > node.Val {
		node.Right = InsertNode(node.Right, element)
	} else if element < node.Val {
		node.Left = InsertNode(node.Left, element)
	}

	return node
}

func main() {
	n := NewNode(10)

	n = InsertNode(n, 20)
	n = InsertNode(n, 30)
	n = InsertNode(n, 8)
	n = InsertNode(n, 3)
	n = InsertNode(n, 13)
	n = InsertNode(n, 19)

	fmt.Println("Inorder traversal : ")
	InorderTraverse(n)

	fmt.Println("\nPreorder traversal : ")
	PreOrderTraverse(n)

	fmt.Println("\nPostorder traversal : ")
	PostOrderTraverse(n)

	exist := NodeExist(n, 19)
	fmt.Println("\nElement 19 exist : ", exist)

	DeleteNode(n, 13)
	fmt.Println("Inorder traversal after deleting element : ")
	InorderTraverse(n)

}
