package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int){
	fmt.Println("Display", number)
}

func TestRunHelloWorld(t *testing.T){
	go RunHelloWorld()
	fmt.Println("TestRunHelloWorld")
	time.Sleep(1 * time.Second)
}

func TestDisplayNumber(t *testing.T){
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}