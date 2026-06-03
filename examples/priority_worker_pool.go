package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Task struct {
	Priority int
	id       int
}

func NewTask(id int, priority int) Task {
	return Task{
		Priority: priority,
		id:       id,
	}
}

func (t Task) Execute() {
	fmt.Println("Executing task : ", t.id, " priority : ", t.Priority)
	time.Sleep(time.Second * 5)
}

type WorkerPool struct {
	wg          sync.WaitGroup
	p1          chan Task
	p2          chan Task
	concurrency int
}

func NewWorkerPool(concurrency int) *WorkerPool {
	return &WorkerPool{
		p1:          make(chan Task, concurrency),
		p2:          make(chan Task, concurrency),
		concurrency: concurrency,
	}
}

func (wp *WorkerPool) Worker(ctx context.Context) {
	defer wp.wg.Done()

	p1 := wp.p1
	p2 := wp.p2

	for p1 != nil || p2 != nil {
		select {
		case <-ctx.Done():
			return

		case t, ok := <-p1:
			if !ok {
				p1 = nil
				continue
			}
			t.Execute()

		case t, ok := <-p2:
			if !ok {
				p2 = nil
				continue
			}
			t.Execute()
		}
	}
}

func (wp *WorkerPool) Start(ctx context.Context) {
	for range wp.concurrency {
		wp.wg.Add(1)
		go wp.Worker(ctx)
	}
}

func (wp *WorkerPool) Submit(t Task) {
	switch t.Priority {
	case 0:
		wp.p1 <- t
	default:
		wp.p2 <- t
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.p1)
	close(wp.p2)
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

func startPriorityScheduler() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wp := NewWorkerPool(runtime.NumCPU())
	wp.Start(ctx)

	go func() {
		for i := 0; i <= 30; i++ {
			num := i + 1
			priority := 1
			if num%3 == 0 && num%5 == 0 {
				priority = 0
			}

			task := NewTask(num, priority)
			wp.Submit(task)
		}

		wp.Stop()
	}()

	wp.Wait()

	fmt.Println("All tasks completed")
}
