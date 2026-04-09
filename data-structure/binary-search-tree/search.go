package main

func NodeExist(node *Node, element int) bool {
	if node == nil {
		return false
	}

	if node.Val == element {
		return true
	}

	if element > node.Val {
		return NodeExist(node.Right, element)
	}

	if element < node.Val {
		return NodeExist(node.Left, element)
	}

	return false
}
