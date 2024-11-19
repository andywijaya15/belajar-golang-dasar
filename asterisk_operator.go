package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	adress1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	var adress2 *Address = &adress1

	adress2.City = "Bandung"
	fmt.Println(adress1) //ikut berubah
	fmt.Println(adress2) //berubah jadi bandung

	// karna type data adress2 pointer menuju Adress jadi tambahkan & didepannya
	// adress2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(adress1)
	// fmt.Println(adress2)

	// asterisk operator
	// siapapun yg menggunakan adress2 maka semuanya akan mengacu ke data adress2
	*adress2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(adress1)
	fmt.Println(adress2)
}
