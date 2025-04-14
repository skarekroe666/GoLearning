package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go count(i)
	}

	time.Sleep(500 * time.Millisecond)
}
func count(i int) {
	fmt.Println(i)
}
