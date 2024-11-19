package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func belajar_function_return_value() {
	result := getHello("Andy")
	fmt.Println(result)
}
