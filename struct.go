package main

import "fmt"

// template / prototype untuk data
// biasanya nama attributenya depannya huruf besar atau pascalcase
type Customer struct {
	Name, Address string
	Age           int
}

// struct method
// function sayHello cuma bisa dipanggil kalau sudah membuat object dari si Customer
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var andy Customer
	andy.Name = "Andy"
	andy.Address = "Wijaya"
	andy.Age = 27
	fmt.Println(andy)
	// atau dengan cara
	lidia := Customer{
		Name:    "Andy",
		Address: "Wijaya",
		Age:     27,
	}
	fmt.Println(lidia)
	// membuat object dari Customer
	budi := Customer{"Andy", "Wijaya", 27}
	fmt.Println(budi)
	budi.sayHello("Joko")
	// representadi data dari aplikasi kita buat dalam bentuk struct
}
