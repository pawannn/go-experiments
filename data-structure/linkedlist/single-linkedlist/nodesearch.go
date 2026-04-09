package main

// Time complexity for search A node:
// Best case: The node will be found at index position 0 -> O(1)
// Worst case: The node will be found at last position -> O(n)

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
