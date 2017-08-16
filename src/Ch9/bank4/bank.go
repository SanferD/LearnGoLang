package bank

import "sync"

var (
	mu sync.RWMutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Withdrawal(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	if amount > balance {
		return false
	}
	deposit(-amount)
	return true
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func deposit(amount int) {
	balance += amount
}
