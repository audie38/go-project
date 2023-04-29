package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (cust Customer) greetCustomer(){
	fmt.Println("Welcome", cust.Name)
}


func main() {
	var cust Customer
	cust.Name = "Audie"
	cust.Address = "Jalan Test"
	cust.Age = 23

	fmt.Println(cust)

	cust1 := Customer{
		Name: "Milson",
		Address: "Jalan Test 1",
		Age: 23,
	}

	fmt.Println(cust1)

	cust.greetCustomer()
	cust1.greetCustomer()
}