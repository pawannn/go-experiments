package main

func (ll *LinkedList) SearchNode(element int) int {
	curr := ll.Head

	i := 0
	for curr != nil {
		if curr.Val == element {
			return i
		}

		curr = curr.Next
		i++
	}

	return -1
}
