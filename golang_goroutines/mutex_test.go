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

type UserBalance struct{
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock(){
	user.Mutex.Lock()
}

func (user *UserBalance) UnLock(){
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int){
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int){
	user1.Lock()
	fmt.Println("Lock User1", user1.Name)
	user1.Change(-amount)

	user2.Lock()
	fmt.Println("Lock User2", user2.Name)
	user2.Change(amount)

	user1.UnLock()
	user2.UnLock()
}

func TestTransfer(t *testing.T){
	user1 := UserBalance{
		Name: "Audie",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name: "Milson",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 200000)
	go Transfer(&user2, &user1, 50000)

	time.Sleep(5 * time.Second)

	fmt.Println(user1)
	fmt.Println(user2)
}