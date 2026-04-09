package main

// Time complexity for deleting A node:
// Best case: The node is at index position 0 -> O(1)
// Worst case: The node is at last position -> O(n)

func (ll *LinkedList) DeleteNode(element int) {
	curr := ll.Head

	if curr.Val == element {
		ll.Head = curr.Next
		return
	}

	for curr != nil && curr.Next != nil {
		if curr.Next.Val == element {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
}
