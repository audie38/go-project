package main

import (
	"fmt"
	"os"
)

func main(){

	// package OS
	fmt.Println(os.Args)
	hostName, err := os.Hostname()
	if err != nil{
		fmt.Println("Error: ", err.Error())
	}else{
		fmt.Println("HostName: ", hostName)
	}

	fmt.Println(os.Getenv("GOPATH"))

}