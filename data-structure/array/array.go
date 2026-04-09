package main

import "fmt"

type Array struct {
	capacity int
	curr     int
	arr      []int
}

func NewArray(size int) *Array {
	return &Array{
		capacity: size,
		curr:     -1,
		arr:      make([]int, size),
	}
}

// Time Complexity to append an array
// Best case: We know the current position of end pointer: O(1)
// Worst case: We know the current position of end pointer: O(1)

func (a *Array) Push(element int) {
	if a.curr == a.capacity-1 {
		fmt.Println("Array is at full capacity")
		return
	}

	a.curr++
	a.arr[a.curr] = element
}

// // Time Complexity to pop element out of an array
// Best case: We know the current position of end pointer: O(1)
// Worst case: We know the current position of end pointer: O(1)

func (a *Array) Pop() int {
	if a.curr == -1 {
		fmt.Println("Array is at empty")
		return -1
	}

	element := a.arr[a.curr]
	a.arr[a.curr] = -1
	a.curr--

	return element
}

// Time Complexity to append an element
// Best case: The element has to insert at end: O(1)
// Worst case: The element has to insert in middle or at start, all the other elements have to shift right to one index: O(n)

func (a *Array) Insert(element int, pos int) {
	if a.curr == a.capacity-1 {
		fmt.Println("Array is at full capacity")
		return
	}

	for i := a.curr; i >= pos; i-- {
		a.arr[i+1] = a.arr[i]
	}

	a.arr[pos] = element
	a.curr++
}

// Time Complexity to delete an element
// Best case: the array has only one element : O(1)
// Worst case: the array has multiple elements, so the array has to be traversed either for search or for shifting: O(n)

func (a *Array) Delete(element int) {
	if a.curr == -1 {
		fmt.Println("Array is at empty")
		return
	}

	pos := -1
	for i := 0; i <= a.curr; i++ {
		if a.arr[i] == element {
			pos = i
		}
	}

	if pos == -1 {
		return
	}

	for i := pos; i < a.curr; i++ {
		a.arr[i] = a.arr[i+1]
	}

	a.curr--
}

func (a *Array) Print() {
	for i := range a.curr + 1 {
		fmt.Println(a.arr[i])
	}
}

func main() {
	a := NewArray(10)

	a.Push(1)
	a.Push(2)
	a.Push(3)
	a.Push(4)
	a.Push(5)
	a.Push(6)
	a.Push(6)
	a.Push(6)
	a.Push(6)

	a.Insert(10, 7)

	a.Pop()

	a.Delete(10)

	a.Print()

}
