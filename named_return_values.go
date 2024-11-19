package main

import "fmt"

// deklarasi di return value nya jadi tidak perlu menggunakan := divariablenya,sehingga jika tidak ada isi dalam named variable nya maka akan mereturn default valuenya
func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Lidia"
	middleName = "Tri"
	// lastName = "Wulansari"
	return firstName, middleName, lastName
}

func belajar_named_return_values() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
