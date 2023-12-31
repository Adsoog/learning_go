package main

import "fmt"

func main() {

	resta := sub(7, 8)
	fmt.Println(resta)
	fmt.Println(sub(5, 5))
}

func sub(x int, y int) int {
	return x - y
}
