package main

// Time complexity of Searching a node
// Best case: The element is at head: O(1)
// Worst case: The element is not at the head: O(n)

func (cl *CircularLinkedList) NodeExist(element int) bool {
	if cl.Head == nil {
		return false
	}

	curr := cl.Head

	for {
		if curr.Val == element {
			return true
		}

		curr = curr.Next
		if curr == cl.Head {
			break
		}
	}

	return false
}
