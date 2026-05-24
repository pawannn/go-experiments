package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

type Task struct {
	id int
}

func (t Task) Execute() {
	fmt.Println(t.id)
	time.Sleep(time.Second * 2)
}

type WorkerPool struct {
	wg sync.WaitGroup

	taskChan    chan Task
	concurrency int
}

func NewWorkerPool(concurrency int) *WorkerPool {
	return &WorkerPool{
		taskChan:    make(chan Task, concurrency),
		concurrency: concurrency,
	}
}

func (wP *WorkerPool) Worker(ctx context.Context) {
	defer wP.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case t, ok := <-wP.taskChan:
			if !ok {
				return
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
	wp.taskChan <- t
}

func (wp *WorkerPool) Stop() {
	close(wp.taskChan)
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	concurrency := runtime.NumCPU()
	wp := NewWorkerPool(concurrency)

	wp.Start(ctx)

	go func() {
		for i := range 30 {
			t := Task{i + 1}
			wp.Submit(t)
		}

		wp.Stop()
	}()

	go func() {
		<-sigChan
		fmt.Println("Termination signal recieved")
		cancel()
	}()

	wp.Wait()

	fmt.Println("All Workers stopped")
}
