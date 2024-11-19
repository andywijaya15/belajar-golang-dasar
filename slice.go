package main

import "fmt"

func slice() {
	names := [...]string{"1", "2", "3", "4"}
	slice1 := names[0:2]
	fmt.Println(slice1)

	slice2 := names[1:]
	fmt.Println(slice2)

	slice3 := names[:3]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	//

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Andy"
	newSlice[1] = "Wijaya"
	// error karena kapasitasnya diset hanya 2
	// newSlice[2] = "error"
	fmt.Println(newSlice)
	// tambah slice menggunakan fungsi append
	newSlice2 := append(newSlice, "append")
	fmt.Println(newSlice2)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// beda slice dan array
	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// pilih array atau slice ? kebanyakan golang menggunakan slice
}
