package main

// Time complexity of Inserting a node at head
// Best case: The element is inserted at head: O(1)
// Worst case: The element is inserted at head: O(1)

func (cl *CircularLinkedList) InsertFront(element int) {
	newNode := NewNode(element)

	if cl.Head == nil {
		cl.Head = newNode
		cl.Head.Next = newNode
		cl.Head.Prev = newNode
		return
	}

	firstElement := cl.Head
	prevElement := cl.Head.Prev

	newNode.Next = firstElement
	newNode.Prev = prevElement

	prevElement.Next = newNode
	firstElement.Prev = newNode

	cl.Head = newNode
}
