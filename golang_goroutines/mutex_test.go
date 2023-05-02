package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
Mutex : prevent race condtion by locking the shared variables/data
*/

func TestMutex(t *testing.T){
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++{
		go func(){
			for j:= 1; j<= 100; j++{
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

type BankAccount struct{
	RWMutex sync.RWMutex
	Balance int
}

func (acc *BankAccount) AddBalance(amount int){
	acc.RWMutex.Lock()
	acc.Balance += amount
	acc.RWMutex.Unlock()
}

func (acc *BankAccount) GetBalancce() int{
	acc.RWMutex.RLock()
	balance := acc.Balance
	acc.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T){
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func(){
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalancce())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalancce())
}