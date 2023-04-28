package main

import "fmt"

func main() {

	// integer
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32, nilai64, nilai16)

	// string
	var (
		name = "Audie Milson"
		a = name[0]
		aString = string(a)
	)

	fmt.Println(name, a, aString)
}