package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func task(id int) error {
	time.Sleep(time.Second * 1)
	if id%2 == 0 {
		return fmt.Errorf("task %d failed", id)
	}

	fmt.Println("task completed", id)
	return nil
}

func startTasks() {
	var g errgroup.Group

	for i := range 10 {
		id := i
		g.Go(func() error {
			return task(id)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("All tasks successful")
}
