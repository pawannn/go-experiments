package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Task struct {
	id   int
	name string
}

func (t *Task) ProcessTask(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Printf("task cancelled: %d\n", t.id)
		return
	default:
	}

	fmt.Println("completing task ID:", t.id, "task name:", t.name)

	workTime := time.Millisecond * time.Duration(rand.Intn(90)+10)

	select {
	case <-ctx.Done():
		log.Printf("task interrupted: %d\n", t.id)
	case <-time.After(workTime):
	}
}

type workerPool struct {
	tasks       []Task
	concurrency int
	taskChan    chan Task
	wg          sync.WaitGroup
}

func (wp *workerPool) worker(ctx context.Context, workerID int) {
	log.Printf("worker %d started\n", workerID)

	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d shutting down\n", workerID)
			return

		case task, ok := <-wp.taskChan:
			if !ok {
				log.Printf("worker %d exiting (channel closed)\n", workerID)
				return
			}

			task.ProcessTask(ctx)
			wp.wg.Done()
		}
	}
}

func (wp *workerPool) Run(ctx context.Context) {
	wp.taskChan = make(chan Task, wp.concurrency)

	for i := 0; i < wp.concurrency; i++ {
		go wp.worker(ctx, i)
	}

	go func() {
		defer close(wp.taskChan)

		for _, task := range wp.tasks {
			select {
			case <-ctx.Done():
				log.Println("stopping task producer")
				return

			case wp.taskChan <- task:
				wp.wg.Add(1)
			}
		}
	}()

	wp.wg.Wait()

	log.Println("all tasks completed")
}

func startWorkerPool() {
	// pprof server
	go func() {
		log.Println("pprof running on :6060")
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		sig := <-sigChan
		log.Printf("received signal: %s\n", sig)

		cancel()
	}()

	// create tasks
	n := 100000

	tasks := make([]Task, n)

	for i := 0; i < n; i++ {
		tasks[i] = Task{
			id:   i,
			name: "work" + fmt.Sprint(i),
		}
	}

	wp := workerPool{
		tasks:       tasks,
		concurrency: 10,
	}

	wp.Run(ctx)

	log.Println("worker pool shutdown complete")
}
