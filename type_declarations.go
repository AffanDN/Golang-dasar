package main

import "fmt"

func main() {
	type NoKTP string

	var ktpJoe NoKTP = "1234567890"

	var contoh string = "222222222"

	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpJoe)
	fmt.Println(contoh)
	fmt.Println(contohKtp)
}