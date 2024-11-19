package main

import "fmt"

func belajar_break() {
	// hentikan perulangan dengan break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}
}
