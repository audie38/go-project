package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Package Strings
	var fullName string = "@Audie Milson"
	var days string = "Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday"

	fmt.Println(strings.Trim(fullName, "@"))
	fmt.Println(strings.ToLower(fullName))
	fmt.Println(strings.ToUpper(fullName))
	fmt.Println(strings.Split(days, ","))
	fmt.Println(strings.Contains(days, "Saturday"))
	fmt.Println(strings.Contains(days, "Holiday"))
	fmt.Println(strings.ReplaceAll(days, "day", ""))

	// Package StrConv
	var angka int = 8
	var booleanString string = "false"
	var floatString string = "3.14"
	var intString string = "8989"

	var boolFalse bool = true

	fmt.Println(strconv.Itoa(angka))
	fmt.Println(strconv.ParseBool(booleanString))
	fmt.Println(strconv.ParseFloat(floatString, 10))
	fmt.Println(strconv.ParseInt(intString, 10, 64))

	fmt.Println(strconv.FormatBool(boolFalse))
	fmt.Println(strconv.FormatInt(898, 10))
}