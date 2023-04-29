package main

import "fmt"

func NewMap(name string) map[string]string{
	if name == ""{
		return nil
	}else{
		return map[string]string{
			"name" : name,
		}
	}
}

func main() {
	// nil only can be used in interface, fucntion, map, slice, pointer, and channel data type
	var person map[string]string = nil
	fmt.Println(person)

	var name string = "Test"

	person = NewMap(name)
	fmt.Println(person)
}