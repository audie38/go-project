package helper

import "fmt"

/**
PascalCase -> Public Access Modifier, can be access from external
camelCase -> Private Access Modifier, cannot be accessed from external
*/

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string){
	fmt.Println("Goodbye", name)
}