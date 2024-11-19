package main

import "fmt"

func belajar_for() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// for dengan statement,init statement untuk kondisi dan post statment untuk counter++
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	// for range
	names := []string{
		"Andy",
		"Wijaya",
	}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	// tanpa index dengan variable menganggur (_)
	for _, name := range names {
		fmt.Println(name)
	}
}
