package bank

import "sync"

var (
	balance int
	mu sync.Mutex
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Withdrawal(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	if amount > balance {
		return false
	}
	balance -= amount
	return true
}

func Balance() {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
