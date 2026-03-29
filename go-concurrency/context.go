package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped :", ctx.Err())
			return
		default:
			fmt.Println("Running")
		}
	}
}

func somework() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(time.Second * 1)
	cancel()

	time.Sleep(time.Second * 1)
}
