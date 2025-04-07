package main

import "fmt"

func Demo() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() { chan1 <- 10 }()

	go func() { chan2 <- "bitch" }()

	for range 2 {
		select {
		case receive1 := <-chan1:
			fmt.Println("Received from channel 1:", receive1)
		case receive2 := <-chan2:
			fmt.Println("Received from channel 2:", receive2)
		}
	}
}
