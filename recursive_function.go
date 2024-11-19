package main

import "fmt"

// membuat function yg memanggil dirinya sendiri untuk contoh seperti factorial

// contoh dalam loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func belajar_recursive_function() {
	result := factorialLoop(3)
	fmt.Println(result)
	// hasilnya sama dengan :
	result = 3 * 2 * 1
	fmt.Println(result)
	// menggunakan factorial recursive
	result = factorialRecursive(3)
	fmt.Println(result)

}

// coba menggunakan recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
