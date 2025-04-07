package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		employee(i, &wg)
	}

	wg.Wait()
	fmt.Println("task complete")
}

func employee(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Employee %v working\n", i)
	fmt.Printf("Employee %v finished working\n", i)
}
