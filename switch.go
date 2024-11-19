package main

import "fmt"

func belajar_switch() {
	name := "Andy"

	switch name {
	case "Andy":
		fmt.Println("Hello Andy")
	case "Lidia":
		fmt.Println("Hello Lidia")
	default:
		fmt.Println("Hello Orang Lain")
	}

	// switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name lebih dari 5")
	case false:
		fmt.Println("Name kurang dari 5")
	}

	// switch tanpa kondisi (bukan pengganti if else)
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Name lebih dari 5")
	case length < 5:
		fmt.Println("Name kurang dari 5")
	}
}
