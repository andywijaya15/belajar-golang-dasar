package main

import "fmt"

// interface adalah tipe data yg abstract,tidak memiliki implementasi secara langsung
// interface memiliki definisi" method dan biasanya digunakan sebagai kontrak
// kontrak itu berarti harus ada yg mengimplementasikannya
type HasName interface {
	// harus punya method ini
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// implementasinya dalam bentuk struct
type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// sekarang animal implementasi dari HasName karena memilik method GetName
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	// tidak perlu mengimplementasikan interfacenya secara manual
	person := Person{Name: "Andy"}
	sayHello(person)
	animal := Animal{Name: "Animal"}
	sayHello(animal)
}
