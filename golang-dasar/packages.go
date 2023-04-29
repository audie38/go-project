package main

import (
	"flag"
	"fmt"
	"math"
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

	// package flag
	var host *string = flag.String("host", "localhost", "Put your database host")
	var username *string = flag.String("username", "root", "Put your database username")
	var password *string = flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println(*host, *username, *password)

	// package math
	pi := 3.14
	fmt.Println(math.Round(pi))
	fmt.Println(math.Floor(pi))
	fmt.Println(math.Ceil(pi))
	fmt.Println(math.Max(3.14, 22.7))
	fmt.Println(math.Min(3.14, 22.7))

}