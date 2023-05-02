package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

/**
- Channel used to send/received data in go routines
- Channel must have send and receive, if not will trigger error / deadlock
- Channel by default pass by reference
- Channel <- : send
- <-Channel  : receive
*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) 
	defer close(channel)

	go func(){
		time.Sleep(2* time.Second)
		channel <- "Audie Milson" //Send
		fmt.Println("Data sent to Channel")
	}()

	fmt.Println(<- channel) //Receive
	time.Sleep(5 * time.Second)
}

// Channel as parameter
func GiveMeResponse(channel chan string){
	time.Sleep(2 * time.Second)
	channel <- "Audie"
}

func TestChannelAsParameter(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "Audie"
}

func OnlyOut(channel <-chan string){
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
}