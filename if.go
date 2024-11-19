package main

import "fmt"

func belajar_if() {
	name := "Andy1"

	if name == "Andy" {
		fmt.Println(name)
	} else if name == "Andy1" {
		fmt.Println(name)
	} else {
		fmt.Println("Salah")
	}

	// if dengan short statement
	if length := len(name); length > 5 {
		fmt.Println(length)
	} else {
		fmt.Println("kurang dari 5")
	}

}
