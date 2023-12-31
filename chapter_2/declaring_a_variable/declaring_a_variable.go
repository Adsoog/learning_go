package main

import "fmt"

func main() {
	var smsSendingLimit int
	var costPerSMS float32
	var hasPermission bool
	var username string

	fmt.Printf("El límite es: %v, el costo es: %f y si tiene permiso es: %v. Por último, su nombre es: %s\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
