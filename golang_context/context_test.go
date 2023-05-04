package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T){
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T){
	contextA := context.Background()
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextB, "c", "C")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	
	fmt.Println(contextC.Value("b"))
}

func CreateCounter(ctx context.Context) chan int{
	destination := make(chan int)
	go func(){
		defer close(destination)
		counter := 1
		for{
			select{
			case <- ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T){
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n:= range destination{
		fmt.Println("Counter", n)
		if n == 10{
			break	
		}
	}

	cancel()
	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T){
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	for n:= range destination{
		fmt.Println("Counter", n)
		if n == 10{
			break	
		}
	}
	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T){
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for n:= range destination{
		fmt.Println("Counter", n)
		if n == 10{
			break	
		}
	}
	fmt.Println(runtime.NumGoroutine())
}