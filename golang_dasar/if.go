package main

import "fmt"

func main() {
	var name string = "Audie Milson"

	// Short Statement
	if length := len(name); length > 10{
		fmt.Println("Terlalu panjang")
	}else{
		if name == "Audie" {
			fmt.Println("Hello Audie")
		}else if name == "Milson"{
			fmt.Println("Hello", name)
		}else{
			fmt.Println("Hello ???")
		}
	}


}