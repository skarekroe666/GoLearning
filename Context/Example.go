package main

import (
	"context"
	"fmt"
	"time"
)

func Example() {
	useBackground()
	useTODO()
	useWithCancel()
	useWithTimeout()
	useWithDeadline()
}

func useBackground() {
	ctx := context.Background()
	fmt.Println("[Background] Root context:", ctx)
}

func useTODO() {
	ctx := context.TODO()
	fmt.Println("[TODO] Placeholder context:", ctx)
}

func useWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("[WithCancel] Cancelled:", ctx.Err())
	}
}

func useWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("[WithTimeout] Work completed")
	case <-ctx.Done():
		fmt.Println("[WithTimeout] Timeout:", ctx.Err())
	}
}

func useWithDeadline() {
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("[WithDeadline] Work completed")
	case <-ctx.Done():
		fmt.Println("[WithDeadline] Deadline reached:", ctx.Err())
	}
}
