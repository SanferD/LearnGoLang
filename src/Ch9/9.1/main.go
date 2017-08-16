package main

import (
	"fmt"
	"Ch9/8.1/bank"
)

func main() {
	bank.Deposit(100)
	bank.Deposit(300)
	if !bank.Withdrawal(150) {
		fmt.Println("Not enough money for Withdrawal")
		return
	}

	fmt.Printf("The remaining balance is %d\n", bank.Balance())
	if !bank.Withdrawal(300) {
		fmt.Println("Not enough money for Widthdrawal")
		return
	}
	fmt.Printf("The remaining balance is %d\n", bank.Balance())
}
