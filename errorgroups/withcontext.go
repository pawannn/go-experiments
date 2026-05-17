package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func TaskWithContext(ctx context.Context, id int) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("task cancelled : ", id)
			return nil
		default:
			fmt.Println("Working on task id : ", id)
			time.Sleep(time.Microsecond * 800)

			if id == 2 {
				return fmt.Errorf("Task failed : %d", id)
			}
		}
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	for i := range 10 {
		id := i

		g.Go(func() error {
			return TaskWithContext(ctx, id)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
