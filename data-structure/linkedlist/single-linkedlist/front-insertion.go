package main

func (ll *LinkedList) InsertFront(element int) {
	if ll.Head == nil {
		ll.Head = NewNode(element)
	} else {
		newNode := NewNode(element)
		newNode.Next = ll.Head

		ll.Head = newNode
	}
}
