package main

import "fmt"

func matematika() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d*e
	fmt.Println(c)
	// augmented assignments
	i := 10
	i += 10
	fmt.Println(i)
	i %= 2
	fmt.Println(i)
	j := 1
	j++
	fmt.Println(j)
}
