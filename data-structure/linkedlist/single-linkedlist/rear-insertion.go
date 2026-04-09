package main

// Time complexity for inserting element at the last:
// Best case: The list is empty -> O(1)
// Worst case: The list is not empty and node will be inserted at last -> O(1)

func (ll *LinkedList) InsertRead(element int) {
	if ll.Head == nil {
		ll.Head = NewNode(element)
	} else {
		curr := ll.Head
		for curr.Next != nil {
			curr = curr.Next
		}

		curr.Next = NewNode(element)
	}
}
