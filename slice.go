package main

import "fmt"

func main() {
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
}
