package main

/* Example #29: https://gobyexample.com/non-blocking-channel-operations */

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "Hi"
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity")
	}
}
