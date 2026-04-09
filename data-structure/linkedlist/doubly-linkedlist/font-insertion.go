package main

// Time complexity for adding A node at front:
// Best case: The node will be added at head -> O(1)
// Worst case: The node will be added at head -> O(1)

func (dl *DoublyLinkedList) InsertFront(element int) {
	if dl.Head == nil {
		dl.Head = NewNode(element)
	} else {
		newNode := NewNode(element)

		newNode.Next = dl.Head
		dl.Head.Prev = newNode

		dl.Head = newNode
	}
}
