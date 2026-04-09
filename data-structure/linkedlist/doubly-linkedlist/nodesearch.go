package main

// Time complexity for searching A node:
// Best case: The node will be found at index position 0 -> O(1)
// Worst case: The node will be found at last position -> O(n)

func (dl *DoublyLinkedList) SearchNode(element int) int {
	curr := dl.Head

	i := 0
	for curr != nil {
		if curr.Val == element {
			return i
		}

		i++
		curr = curr.Next
	}

	return -1
}
