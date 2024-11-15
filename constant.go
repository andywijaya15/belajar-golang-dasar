package main

import "fmt"

func constant() {
	const (
		firstName = "Andy"
		lastName  = "Wijaya"
	)
	fmt.Println(firstName, lastName)
}
