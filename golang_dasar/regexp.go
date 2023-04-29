package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`([a-z])`)

	fmt.Println(regex.MatchString("Audie"))
	fmt.Println(regex.MatchString("audie"))
	fmt.Println(regex.FindAllString("Audie Milson Edison", 10))
}