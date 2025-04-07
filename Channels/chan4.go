package main

import (
	"fmt"
	"time"
)

func SelectEx() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(200 * time.Second)
		chan1 <- "hello bitch"
	}()

	go func() {
		time.Sleep(250 * time.Second)
		chan2 <- "yo skarekroe"
	}()

	select {
	case receive1 := <-chan1:
		fmt.Println("Received from channel 1:", receive1)
	case receive2 := <-chan2:
		fmt.Println("Received from channel 2:", receive2)
	}
}
