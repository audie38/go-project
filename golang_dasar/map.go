package main

import "fmt"

func main() {

	// map[key data type]value data type
	person := map[string]string{
		"name":    "Audie",
		"address": "Jakarta",
	}

	fmt.Println(person)
	person["title"] = "Programmer"
	fmt.Println(person, len(person))
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	
	// init map without assigned data
	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Audie Milson"
	book["wrong"] = "Ups"

	fmt.Println("Before", book)
	delete(book, "wrong")
	fmt.Println("After", book)

}