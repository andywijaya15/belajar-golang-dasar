package main

import "fmt"

func main() {
	type NoKtp string
	var ktpAndy NoKtp = "11111111"
	fmt.Println(ktpAndy)
	var contoh string = "22222222"
	var contohKtp NoKtp = NoKtp(contoh)
	fmt.Println(contohKtp)
}
