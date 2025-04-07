package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Incrememnt()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter Value:", counter.Value())
}

type SafeCounter struct {
	count int
	mu    sync.Mutex
}

func (c *SafeCounter) Incrememnt() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
