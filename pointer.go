package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// pass by value,mengcopy variable ke variable lain tapi tidak saling terhubung
	adress1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	adress2 := adress1 //copy value

	adress2.City = "Bandung"
	fmt.Println(adress1) //tidak berubah
	fmt.Println(adress2) //berubah karena duplikat dari address 1 dan perubahannya tidak berimpact ke address1

	// pointer adalah kemampuan membuat reference ke lokasi data di memory yg sama tanpa menduplikasi data yg ada atau bisa disebut dengan pass by reference
	// misal kita tidak mau menduplikasi data yg adas
	// kalau menduplikasi data terus menerus akan menyebabkan boros memory
	// contoh penggunaan pointer
	// ketika data adress3 dirubah,adress 1 juga ikut berubah
	var adress3 *Address = &adress1
	adress3.City = "Semarang"
	fmt.Println(adress1)
	// tanda & didepannya berarti pointer
	fmt.Println(adress3)
}
