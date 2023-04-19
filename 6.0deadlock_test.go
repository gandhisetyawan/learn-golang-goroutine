package learn_golang_goroutine

import (
	"fmt"
	"sync"
)

type UserBalance struct {
	sync.Mutex //Mutex sync.Mutex
	Name       string
	Balance    int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)
}
