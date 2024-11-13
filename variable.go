package main

import "fmt"

func variable() {
	name := "Andy Wijaya"
	fmt.Println(name)

	name = "Andy"
	fmt.Println(name)

	var (
		firstName = "Andy"
		lastName  = "Wijaya"
	)
	fmt.Println(firstName, lastName)
}
