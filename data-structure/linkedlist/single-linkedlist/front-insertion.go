package main

// Time complexity for adding A node at front:
// Best case: The node will be added at head -> O(1)
// Worst case: The node will be added at head -> O(1)

func (ll *LinkedList) InsertFront(element int) {
	if ll.Head == nil {
		ll.Head = NewNode(element)
	} else {
		newNode := NewNode(element)
		newNode.Next = ll.Head

		ll.Head = newNode
	}
}
