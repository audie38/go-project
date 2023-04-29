package main

import (
	"fmt"
	"gopath/database"
	_ "gopath/helper"
)

func main() {
	// helper.SayHello("Audie")
	result := database.GetDataBase()
	fmt.Println(result)
}