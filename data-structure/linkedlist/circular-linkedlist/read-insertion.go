package main

// Time complexity of Inserting a node at rear
// Best case: The element is inserted at rear which is the prev address of head: O(1)
// Worst case: The element is inserted at rear which is the prev address of head: O(1)

func (cl *CircularLinkedList) InsertRear(element int) {
	newNode := NewNode(element)

	if cl.Head == nil {
		cl.Head = newNode
		cl.Head.Next = newNode
		cl.Head.Prev = newNode
		return
	}

	frontElement := cl.Head
	rearElement := cl.Head.Prev

	rearElement.Next = newNode
	frontElement.Prev = newNode

	newNode.Prev = rearElement
	newNode.Next = frontElement
}
