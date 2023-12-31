package main

import "fmt"

func main() {

	sendsSoFar := 430
	const sendsAdd = 25
	fmt.Printf("you've sent %v messages", incrementSends(sendsSoFar, sendsAdd))
}

func incrementSends(sendsSoFar int, sendsAdd int) int {

	sendsSoFar = sendsSoFar + sendsAdd
	return sendsSoFar
}
