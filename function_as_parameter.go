package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func belajar_function_as_parameter() {
	// bisa langsung kirim functionnya
	name := "Anjing"
	sayHelloWithFilter(name, spamFilter)
	// atau bisa function as variable dulu seperti ini
	name = "Andy"
	filter := spamFilter
	sayHelloWithFilter(name, filter)
	// menggunakan type declaration
	sayHelloWithFilterWithTypeDeclarations(name, filter)
}

// di golang gunakanlah type declarations untuk mempermudah penulisan seperti diatas

type Filter func(string) string

// contoh penggunaan :
func sayHelloWithFilterWithTypeDeclarations(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}
