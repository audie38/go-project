package golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/**
- Channel used to send/received data in go routines
- Channel must have send and receive, if not will trigger error / deadlock
- Channel by default pass by reference
- Channel <- : send
- <-Channel  : receive
- Channel by default can only send/receive 1 data
- Buffered : Capacity of how many data can be send/received in channel
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

func TestBufferedChannel(t *testing.T){
	channel := make(chan string, 3)
	defer close(channel)

	go func(){
		channel <- "Audie"
		channel <- "Milson"

		fmt.Println("len(channel)", len(channel))
		fmt.Println("cap(channel)", cap(channel))

		fmt.Println("Channel[0]", <- channel)
		fmt.Println("Channel[1]", <- channel)
	}()
	
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

func TestRangeChannel(t *testing.T){
	channel := make(chan string)

	go func(){
		for i := 0; i<10; i++{
			channel <- "Loop No." + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel{
		fmt.Println("Receiving Data", data)
	}

	fmt.Println("Done")
}