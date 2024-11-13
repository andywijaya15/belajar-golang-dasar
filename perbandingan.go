package main

import "fmt"

func main() {
	name1 := "Andy"
	name2 := "Andy"

	result := name1 == name2
	fmt.Println(result)
	result = name1 != name2
	fmt.Println(result)
}
