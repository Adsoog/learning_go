package main

import "fmt"

func main() {

	firstName, _ := getNames() // And remenber de under score "_" is use to ignore te second part of the function
	fmt.Println("Welcome to mi cave,", firstName)

}

func getNames() (string, string) {
	return "John", "Doe"
}
