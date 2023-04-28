package main

import "fmt"

func main() {
	counter := 1

	for counter <= 7 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Println("")

	for i:= 0; i< 10; i++ {
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("")

	slice := []string{"Audie", "Milson"}

	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	for index, name := range slice{
		fmt.Println("index", index, "=", name)
	}

	for _, name := range slice{
		fmt.Println("nama =", name)
	}
}