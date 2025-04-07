package main

import (
	"fmt"
)

func main() {

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(10)
	// }

	result := make(chan int)

	go sum(result, 1, 2)

	res := <-result
	fmt.Println("Result:", res)

	done := make(chan bool)

	go task(done)

	<-done
	fmt.Println("-----------------------------------------")
	// Example()
	// Demo()
	SelectEx()
	fmt.Println("-----------------------------------------")
}

// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number:", num)
// 		time.Sleep(1 * time.Second)
// 	}
// }

func sum(result chan int, a, b int) {
	result <- a + b
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing...")
}
