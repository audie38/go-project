package main

import (
	"fmt"
	"reflect"
)

func random() interface{} {
	return "OK"
}

func main() {
	var result interface{} = random()
	
	var res string = result.(string)
	fmt.Println(res)

	switch value:= result.(type){
		case string:
			fmt.Println("String", value)	
		case int:
			fmt.Println("Int", value)	
		default:
			fmt.Println(reflect.TypeOf(value), value)	
	}
	
}