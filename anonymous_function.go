package main

import "fmt"

// akan lebih mudah langsung membuat function secara langsung pada variable / parameter tanpa harus membuat function lebih dulu,ini disebut anonymous function

// type declarations
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func belajar_anonymous_function() {
	// bisa diassign langsung ke variable
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Andy", blacklist)
	// atau langsung lempar function
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
