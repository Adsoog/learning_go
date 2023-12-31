package main

// Debemos crear un string que tenga el siguiente megnsaje "hi NAME, your open rate is OPENRATE"

import "fmt"

func main() {

	const name = "Saul Goodman"
	const openRate = 30.5

	// como hacemos para asignarle variables y string a una variable? usamos la funcion Sprint de fmt

	msg := fmt.Sprintf(
		"hi %s, your open rate is %f ",
		name,
		openRate,
	)

	fmt.Println(msg)
}
