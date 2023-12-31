package main

import "fmt"

func main() {

	michis := 2
	if michis < 3 {
		fmt.Println("Tiene 2 michis")
	}

	if dogos := 9; dogos < 8 {
		fmt.Println("Tienes pocos dogos")
	} else {
		fmt.Println("Tienes muchos dogos")
	}

}

// if INITIAL_STATEMENT; CONDITION {}
