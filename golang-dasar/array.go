package main

import "fmt"

func main() {
	var names [2]string

	names[0] = "Audie"
	names[1] = "Milson"

	fmt.Println(names, names[0], names[1])

	var values = [3]int{
		90, 95, 80,
	}

	fmt.Println(values, len(values))

	values[2] = 100
	fmt.Println(values, len(values))
}