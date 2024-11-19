package main

import "fmt"

// kemampuan sebuah class berinteraksi dengan data disekitarnya

// hati hati menggunakan closure karena bisa merubah variable tanpa kita ketahui dalam 1 code yg panjang

func belajar_closure() {
	counter := 0
	// function disini bisa mengakses data diatasnya dana memodifikasinya
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	increment()
	increment()
	fmt.Println(counter)
}
