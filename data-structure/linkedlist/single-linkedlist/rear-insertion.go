package main

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
