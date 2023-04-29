package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hn HasName) {
	fmt.Println("Hello", hn.GetName())
}

type Person struct{
	Name string
}

type Animal struct{
	Name string
}

func (person Person) GetName() string{
	return person.Name
}

func (animal Animal) GetName() string{
	return animal.Name
}

func emptyInterface(i int) interface{}{
	if i == 1{
		return 1
	}else if i == 2{
		return true
	}else {
		return "empty"
	}
}

func main() {
	var milson Person
	milson.Name = "Audie"
	
	SayHello(milson)

	cat := Animal{
		Name: "Kucing",
	}

	SayHello(cat)

	var empty interface{} = emptyInterface(3)

	fmt.Println(empty)
}