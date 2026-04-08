package main

import "fmt"

type Queue struct {
	Front    int
	Rear     int
	capacity int
	arr      []int16
}

func NewQueue(size int) *Queue {
	return &Queue{
		Front:    0,
		Rear:     0,
		capacity: size,
		arr:      make([]int16, size),
	}
}

func (q *Queue) Enqueue(element int16) {
	if q.Rear == q.capacity {
		fmt.Println("Queue is full")
	} else {
		q.arr[q.Rear] = element
		q.Rear++
	}
}

func (q *Queue) Dequeue() int16 {
	if q.Front == q.Rear {
		fmt.Println("Queue is empty")
	} else {
		element := q.arr[q.Front]
		q.Front++

		return element
	}

	return -1
}

func (q *Queue) PrintQueue() {
	for i := q.Front; i < q.Rear; i++ {
		fmt.Println(q.arr[i])
	}
}

func main() {
	q := NewQueue(5)

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)

	fmt.Println("After elements enqueue")
	q.PrintQueue()

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	fmt.Println("After elements dequeue")
	q.PrintQueue()
}
