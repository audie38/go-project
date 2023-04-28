package main

import "fmt"

func main() {
	type NoKTP string

	var testKtp NoKTP = "1234567890"
	fmt.Println(testKtp)
	fmt.Println(NoKTP("0987654321"))
}