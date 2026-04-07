package main

import "fmt"

// Time Complexity of Adding an element : O(1)
// Time Complexity of Removing an element : O(1)
// Time Complexity of Finding an element : O(N)

type Stack struct {
	top      int
	capacity int
	array    []int16
}

func NewStack(size int) *Stack {
	return &Stack{
		top:      0,
		capacity: size,
		array:    make([]int16, size),
	}
}

func (s *Stack) Append(element int16) {
	if s.top == s.capacity {
		fmt.Println("Stack is at full capacity")
		return
	}

	s.array[s.top] = element
	s.top++
}

func (s *Stack) Pop() int16 {
	if s.top == 0 {
		fmt.Println("Stack Empty")
		return -1
	}

	s.top--

	lastElement := s.array[s.top]

	s.array[s.top] = 0

	return lastElement
}

func (s *Stack) Find(element int16) int {
	if s.top == 0 {
		return -1
	}

	for i := range s.top {
		if s.array[i] == element {
			return i
		}
	}

	return -1
}

func (s Stack) PrintStack() {
	if s.top == 0 {
		return
	}

	for i := range s.top {
		fmt.Printf("%d ", s.array[i])
	}
	fmt.Println()
}

func main() {
	s := NewStack(4)

	s.Append(10) // 10
	s.Append(20) // 10, 20
	s.Append(30) // 10, 20, 30
	s.Append(40) // 10, 20, 30, 40

	s.PrintStack() // 10, 20, 30, 40

	idx := s.Find(30)
	fmt.Printf("element %d is at index %d\n", 30, idx)

	i := s.Pop()
	fmt.Println("Popped element : ", i)

	i = s.Pop()
	fmt.Println("Popped element : ", i)

	i = s.Pop()
	fmt.Println("Popped element : ", i)

	i = s.Pop()
	fmt.Println("Popped element : ", i)

}
