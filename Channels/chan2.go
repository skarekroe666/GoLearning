package main

import (
	"fmt"
	"time"
)

func Example() {

	emailChan := make(chan string, 10)
	done := make(chan bool)

	go sendEmail(emailChan, done)

	for i := 1; i <= 50; i++ {
		emailChan <- fmt.Sprintf("mail%d@fake.com", i)
	}

	close(emailChan) //THIS IS IMPORTANT
	<-done
}

func sendEmail(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for v := range emailChan {
		fmt.Println("Sending email to:", v)
		time.Sleep(100 * time.Millisecond)
	}
}
