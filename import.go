package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Andy")
	fmt.Println(result)
	// bisa mengakses variable Application karena access modifier membolehkan untuk mengakses variable dengan huruf besar diluar package tersebut
	fmt.Println(helper.Application)
	// tidak bisa diakses karena menggunakan huruf kecil
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBye)
}
