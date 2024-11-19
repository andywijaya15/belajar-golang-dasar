package main

import "fmt"

// contoh panic untuk menghentikan program dan recover untuk menjalankan program kembali

func endApp() {
	fmt.Println("End App")
	// jalankan recover di defer function untuk memastikan proses berjalan kembali setelah panic
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func runApp(error bool) {
	// defer function tetap berjalan meskipun aplikasi dijadikan panic yg menghentikan apps
	defer endApp()
	if error {
		// program akan langsung dihentikan ketika panic dan kode dibawahnya tidak dieksekusi lagi
		panic("Ups Error")
	}
	// jangan gunakan recover disini karena tidak akan dieksekusi ketika panic berjalan
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func main() {
	runApp(true)
}
