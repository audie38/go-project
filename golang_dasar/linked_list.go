package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	// Double Linked List
	data := list.New()
	data.PushBack("Test")
	data.PushBack("Test1")
	data.PushBack("Test2")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	fmt.Println("Loop Access Linked List")
	
	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("Reverse Loop Access Linked List")

	for i := data.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

	// Circular List
	fmt.Println("Circular List")

	var data1 *ring.Ring = ring.New(5)
	for i := 0; i< data.Len(); i++{
		data1.Value = "Value" + strconv.FormatInt(int64(i), 10)
		data1 = data1.Next()
	}

	data1.Do(func(value interface{}){
		fmt.Println(value)
	})
}