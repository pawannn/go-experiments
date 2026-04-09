package main

// Time complexity for deleting A node:
// Best case: The node is at index position 0 -> O(1)
// Worst case: The node is at last position -> O(n)

func (dl *DoublyLinkedList) DeleteNode(element int) {
	if dl.Head == nil {
		return
	}

	// If First element is the element to delete
	if dl.Head.Val == element {
		dl.Head = dl.Head.Next
		if dl.Head != nil {
			dl.Head.Prev = nil
		}

		return
	}

	curr := dl.Head
	for curr.Next != nil {
		// if the next element if the target element
		if curr.Next.Val == element {
			toDelete := curr.Next

			// the current element points to next element of deleted element
			curr.Next = toDelete.Next

			// if the delete element is not the last element
			if toDelete.Next != nil {
				// point the next element prev to current element
				toDelete.Next.Prev = curr
			}
			return
		}

		curr = curr.Next
	}
}
