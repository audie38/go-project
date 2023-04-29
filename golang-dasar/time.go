package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())

	utc := time.Date(2023, time.April, 29, 23, 7, 0, 0, time.UTC)
	fmt.Println("UTC Local: ", utc.Local())

	parse, _ := time.Parse("2006-01-02", "2023-04-29")
	fmt.Println(parse)

}