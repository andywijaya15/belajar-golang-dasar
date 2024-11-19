package main

import "fmt"

// walaupun method menempel didalam struct,tapi sebenarnya data struct yg diakses dalam method adalah pass by value
// sangat direkomendasikan menggunakan pointer di method sehingga tidak boros memory karena terlalu banyak duplikasi data ketika memanggil method

type Man struct {
	Name string
}

// karena disini parameternya adalah value,karena value artinya diduplikasi datanya sehingga data di variable andy tidak berubah
// func (man Man) Married() {
// 	man.Name = "Mr. " + man.Name
// }

// untuk mengubah variable andy ubah type structnya menjadi pointer
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	andy := Man{"Andy"}
	andy.Married()
	fmt.Println(andy.Name)
}
