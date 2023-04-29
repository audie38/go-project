package main

import (
	"fmt"
	"strings"
)

// Non return function
func sayHello() {
	fmt.Println("Hello")
}

// function with parameters
func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

// function with return values
func getHello(name string) string{
	if name == ""{
		return "Hello ???"
	}else{
		return "Hello" + name
	}
}

func add(a int, b int) int{
	return a + b
}

func substract(a int, b int) int{
	return a - b
}

func multiply(a int, b int) int{
	return a * b
}

func divide(a int, b int) int{
	return a / b
}

// function returning multiple values
func getFullName() (string, string){
	return "Audie", "Milson"
}

// named return values
func getCompleteName() (firstName, lastName string){
	firstName = "Milson"
	lastName = "Audie"

	return
}

// variadic function
func sumAll(numbers ...int) (total int){
	total = 0

	for _, number := range numbers{
		total +=  number
	}

	return
}

// function value
func getGoodBye(name string) string{
	return "Good Bye " + name
}

// function as parameter
type Filter func(string)string

func sayHelloWithFilter(name string, filter Filter){
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string{
	if name == "Bad"{
		return "..."
	}else{
		return name
	}
}

func upperFilter(name string)string{
	return strings.ToUpper(name)
}

type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You are Blocked", name)
	}else{
		fmt.Println("Welcome", name)
	}
}

// recursive
func factorialLoop(value int) int{
	if(value == 1){
		return 1
	}else{
		return value * factorialLoop(value - 1)
	}
}


func main() {
	sayHello()
	sayHelloTo("Audie", "Milson")

	var name string
	fmt.Println(getHello(name))
	
	fmt.Println(add(5, 5))
	fmt.Println(substract(5, 5))
	fmt.Println(multiply(5, 5))
	fmt.Println(divide(5, 5))

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	firstName, lastName = getCompleteName()
	fmt.Println(firstName, lastName)

	fmt.Println(sumAll(10, 10, 10, 10, 10))

	numbers := []int{15, 15, 15, 15}
	fmt.Println(sumAll(numbers...))

	goodBye := getGoodBye
	fmt.Println(goodBye("Test"))

	name = "Bad"
	sayHelloWithFilter(name, spamFilter)
	sayHelloWithFilter(name, upperFilter)

	// anonymous function
	blackList := func(name string) bool{
		return name == "Bad"
	}

	registerUser("Admin", blackList)
	registerUser("Bad", func(name string) bool{
		return name == "Bad"
	})

	fmt.Println(factorialLoop(5))
}