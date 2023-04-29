package main

import "fmt"

func main() {
	const name string = "Audie"
	const age = 23

	fmt.Println(name)
	fmt.Println(age)

	const (
		firstName string = "Audie"
		lastName = "Milson"
	)

	fmt.Println(firstName, lastName)
}