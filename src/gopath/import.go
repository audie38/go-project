package main

import "gopath/helper"

func main() {
	helper.SayHello("Audie")
	// helper.sayGoodBye("Audie") // Error because private access modifier
}