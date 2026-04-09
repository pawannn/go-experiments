package main

func (ll *LinkedList) DeleteNode(element int) {
	curr := ll.Head

	if curr.Val == element {
		ll.Head = curr.Next
		return
	}

	for curr != nil && curr.Next != nil {
		if curr.Next.Val == element {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
}
