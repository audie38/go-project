package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"20"`
}

func main() {
	sample := Sample{"Audie"}

	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0))
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
}