package main

import "fmt"

// number adalah slice integer (lebih simple dari yg bawah sendiri)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func belajar_variadic_function() {
	// jadi tidak perlu membuat slice dulu
	fmt.Println(sumAll(10, 10, 10))
	// tidak perlu seperti ini
	fmt.Println(sumAllSlice([]int{10, 10}))
	// jika memang sudah terlanjur memiliki slice yg mau dikirim ke variadic_function maka tambahkan ... diakhir variable supaya tidak hanya dibaca sebagai 1 data,contohnya :
	terlanjurSlice := []int{10, 10}
	total := sumAll(terlanjurSlice...)
	fmt.Println(total)
	// caranya begitu melempar slice ke variadic function
}

// bayangkan kalau tipenya slice
func sumAllSlice(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
