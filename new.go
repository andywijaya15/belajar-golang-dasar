package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// untuk membuat pointer dengan data kosong gunakan new func
	var adress1 *Address = new(Address)
	var adress2 *Address = adress1

	adress2.City = "Bandung"
	fmt.Println(adress1)
	fmt.Println(adress2)
}
