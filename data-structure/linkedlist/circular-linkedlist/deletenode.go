package main

// Time complexity of deleting a node
// Best case: The element to delete is at head: O(1)
// Worst case: The element is in between somewhere in the list: O(n)

func (cl *CircularLinkedList) DeleteNode(element int) {
	if cl.Head == nil {
		return
	}

	curr := cl.Head

	for {
		if curr.Val == element {
			if curr.Next == curr {
				cl.Head = nil
				return
			}

			curr.Prev.Next = curr.Next
			curr.Next.Prev = curr.Prev

			if curr == cl.Head {
				cl.Head = curr.Next
			}

			return
		}

		curr = curr.Next
		if curr == cl.Head {
			break
		}
	}
}
