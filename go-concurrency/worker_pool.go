package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Task struct {
	id   int
	name string
}

func (t *Task) ProcessTask() {
	fmt.Println("completing task ID : ", t.id, " task name : ", t.name)
	time.Sleep(time.Second * 2)
}

type workerPool struct {
	tasks       []Task
	concurrency int
	taskChan    chan Task
	wg          sync.WaitGroup
}

func (wp *workerPool) worker() {
	for task := range wp.taskChan {
		task.ProcessTask()
		wp.wg.Done()
	}
}

func (wp *workerPool) Run() {
	wp.wg.Add(len(wp.tasks))
	wp.taskChan = make(chan Task, len(wp.tasks))
	for range wp.concurrency {
		go wp.worker()
	}

	for _, task := range wp.tasks {
		wp.taskChan <- task
	}

	wp.wg.Wait()
	close(wp.taskChan)
}

func startWorkerPool() {
	tasks := make([]Task, 20)
	for i := range 20 {
		tasks[i].id = i
		tasks[i].name = "work" + fmt.Sprint(i)
	}

	wp := workerPool{
		tasks:       tasks,
		concurrency: runtime.NumCPU(),
	}

	wp.Run()
}
