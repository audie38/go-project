package main

import "fmt"

func main() {
	var name string = "Audie"

	switch name {
	case "Audie":
		fmt.Println("Hello Audie")
	case "Milson":
		fmt.Println("Hello Milson")
	default:
		fmt.Println("Hello ???")
	}

	// Short Statement
	switch length := len(name); length > 5{
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Sudah sesuai")
	}

	// Switch without condition
	name = "Audie Milson"
	length := len(name)
	switch{
	case length > 10:
		fmt.Println("Nama > 10 karakter")
	case length > 5:
		fmt.Println("Nama > 5 karakter")
	default:
		fmt.Println("Sudah sesuai")
	}
}