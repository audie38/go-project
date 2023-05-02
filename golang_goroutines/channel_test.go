package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

/**
- Channel used to send/received data in go routines
- Channel must have send and receive, if not will trigger error / deadlock
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