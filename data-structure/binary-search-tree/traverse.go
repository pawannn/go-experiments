package main

import "fmt"

func InorderTraverse(root *Node) {
	if root == nil {
		return
	}

	InorderTraverse(root.Left)
	fmt.Printf("%d ", root.Val)
	InorderTraverse(root.Right)
}

func PreOrderTraverse(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Val)
	InorderTraverse(root.Left)
	InorderTraverse(root.Right)
}

func PostOrderTraverse(root *Node) {
	if root == nil {
		return
	}

	InorderTraverse(root.Left)
	InorderTraverse(root.Right)
	fmt.Printf("%d ", root.Val)
}
