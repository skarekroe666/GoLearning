package main

import (
	"fmt"
	"time"
)

func main() {
	go count()
	count()
}
func count() {
	for i := 1; i < 6; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(i)
	}
}
