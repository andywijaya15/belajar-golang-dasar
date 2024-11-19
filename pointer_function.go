package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// func ChangeCountryToIndonesia(address Address) {
// 	address.Country = "Indonesia"
// }

// untuk menerima parameter pointer
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// saat dikirimkan seperti ini maka variable adress tidak akan berubah karena pass by value
	// adress1 := Address{}
	// ChangeCountryToIndonesia(adress1)
	// kirimkan pointernya untuk mengubah langsung variablenya
	adress1 := &Address{}
	ChangeCountryToIndonesia(adress1)
	fmt.Println(adress1)
	// kalau terlanjur variablenya bukan pointer tinggal tambahkan & saja di parameter saat pemanggilan function untuk mengirim pointernya
	adress2 := Address{}
	ChangeCountryToIndonesia(&adress2)
	fmt.Println(adress2)
}
