package main

import "fmt"

// memanggil function setelah function selesai dieksekusi

func logging() {
	fmt.Println("Selesai memanggil function")
}

func logging2() {
	fmt.Println("Selesai memanggil function 2")
}

func logging3() {
	fmt.Println("Selesai memanggil function 3")
}

func runApplication() {
	// ketika ada error di line ini function dibawahnya tidak dieksekusi
	// fmt.Println("Run Application")
	// panic("Error")
	// logging()
	// lebih baik gunakan defer
	defer logging()
	defer logging2()
	defer logging3()
	fmt.Println("Run Application")
	// panic("Error")
}

func main() {
	runApplication()
}
