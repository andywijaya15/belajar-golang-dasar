package helper

import "fmt"

var version = "1.0.0" // access modifier di golang menggunakan huruf besar dan kecil variable apapun itu,jika variable diawali huruf besar maka bisa diakses diluar package ini,jika diawali huruf kecil maka hanya untuk package ini
var Application = "golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

// huruf depan harus besar
func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Andy")
	fmt.Println(version)
}
