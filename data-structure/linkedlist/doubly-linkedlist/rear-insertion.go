package main

// Time complexity for inserting element at the last:
// Best case: The list is empty -> O(1)
// Worst case: The list is not empty and node will be inserted at last -> O(n)

func (dl *DoublyLinkedList) InsertRear(element int) {
	if dl.Head == nil {
		dl.Head = NewNode(element)
	} else {
		curr := dl.Head
		for curr.Next != nil {
			curr = curr.Next
		}

		newNode := NewNode(element)
		curr.Next = newNode
		newNode.Prev = curr
	}
}
