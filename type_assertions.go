package main

import "fmt"

func random() any {
	// karena disini string jadi bisa di assertions ke string
	return "OK"
}

func main() {
	result := random()
	// mengkonversi dari type any ke type yg kita tau
	resultString := result.(string)
	fmt.Println(resultString)
	// jangan sampai salah type yg akan menyebabkan panic
	// resultInt := result.(int)
	// fmt.Println(resultInt)
	// agar lebih aman kita menggunakan switch expression untuk melakukan type assertionsnya
	cobaResult := random()
	switch value := cobaResult.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknwon", value)
	}
}
