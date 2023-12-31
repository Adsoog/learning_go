package main

import "fmt"

func main() {

	messageLen := 10
	maxMessageLen := 21
	fmt.Println(
		"Trying to send a message of length:",
		messageLen,
	)

	if messageLen < maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}
