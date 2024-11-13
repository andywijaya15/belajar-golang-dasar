package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Andy"
	names[1] = "Wijaya"
	fmt.Println(names[0], names[1])
	newNames := [3]int{10, 10, 10}
	fmt.Println(newNames)
	autoArrayDefine := [...]int{10, 10, 10}
	fmt.Println(autoArrayDefine)
}
