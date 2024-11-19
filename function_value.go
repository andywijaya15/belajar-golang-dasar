package main

import "fmt"

// menyimpan function dalam variable
func getGoodBye(name string) string {
	return "Good Bye " + name
}

func belajar_function_value() {
	// menggunakan function menjadi value
	goodbye := getGoodBye
	fmt.Println(goodbye("Andy"))
}
