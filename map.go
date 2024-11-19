package main

import "fmt"

func belajar_map() {
	person := map[string]string{
		"name":    "Andy",
		"address": "Semarang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	// kalo akses map yg tidak ada maka tampil default valuenya sesuai tipenya
	fmt.Println(person["names"])

	book := make(map[string]string)
	book["title"] = "Andy"
	book["author"] = "Andy"

	fmt.Println(book)

	delete(book, "author")
	fmt.Println(book)
}
