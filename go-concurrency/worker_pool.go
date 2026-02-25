package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

type Task struct {
	id   int
	name string
}

func (t *Task) ProcessTask() {
	fmt.Println("completing task ID : ", t.id, " task name : ", t.name)
	workTime := time.Millisecond * time.Duration(rand.Intn(90)+10)
	time.Sleep(workTime)
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
	wp.taskChan = make(chan Task, wp.concurrency)
	for range wp.concurrency {
		go wp.worker()
	}

	for _, task := range wp.tasks {
		wp.wg.Add(1)
		wp.taskChan <- task
	}

	close(wp.taskChan)
	wp.wg.Wait()
}

func startWorkerPool() {
	go func() {
		log.Println("pprof running on :6060")
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	n := 100000
	tasks := make([]Task, n)
	for i := range n {
		tasks[i].id = i
		tasks[i].name = "work" + fmt.Sprint(i)
	}

	wp := workerPool{
		tasks:       tasks,
		concurrency: 10,
	}

	wp.Run()
}
