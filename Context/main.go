package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go doWork(ctx)

	time.Sleep(2 * time.Second)
	cancel() // cancel the context

	time.Sleep(1 * time.Second)

	fmt.Println("--------------------------------------------------")
	Example()
	fmt.Println("--------------------------------------------------")
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
