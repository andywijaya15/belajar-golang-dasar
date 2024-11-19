package main

import "fmt"

// tipe data nil / null cuma bisa digunakan di type data interface,function,map,slice,pointer dan channel diluar ini akan menggunakan default value

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Andy")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data["name"])
	}
}
