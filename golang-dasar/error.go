package main

import (
	"errors"
	"fmt"
)

func division(val int, div int) (int, error) {
	if div == 0 {
		return 0, errors.New("Divided with 0")
	}else{
		return val/div, nil
	}
}

func main() {
	res, err := division(5, 0)

	if(err != nil){
		fmt.Println("Error: ", err.Error())
	}else{
		fmt.Println("Result: ", res)
	}
}