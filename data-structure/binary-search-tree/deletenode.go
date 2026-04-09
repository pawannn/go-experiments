package main

func getRightMin(node *Node) int {
	for node.Left != nil {
		node = node.Left
	}

	return node.Val
}

func DeleteNode(node *Node, element int) *Node {
	if node == nil {
		return nil
	}

	if element > node.Val {
		node.Right = DeleteNode(node.Right, element)
	} else if element < node.Val {
		node.Left = DeleteNode(node.Left, element)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			rightmin := getRightMin(node.Right)
			node.Val = rightmin
			node.Right = DeleteNode(node.Left, rightmin)
		}
	}

	return node
}
