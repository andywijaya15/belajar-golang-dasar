package main

import "fmt"

func getFullName() (string, string) {
	return "Andy", "Wijaya"
}

func belajar_return_multiple_value() {
	// wajib ditangkap semua valuenya
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	// menghiraukan salah 1 return value
	firstName, _ = getFullName()
	fmt.Println(firstName)
}
