package bank

var balances = make(chan int)
var deposits = make(chan int)
var withdrawals = make(chan withdrawalStruct)

type validCh chan bool

type withdrawalStruct struct {
	isValid validCh
	amount int
}

func Deposit(amount int) {
	deposits <- amount
}

func Withdrawal(amt int) bool {
	myWithdraw := withdrawalStruct{ make(validCh), amt }
	withdrawals <- myWithdraw
	return <- myWithdraw.isValid
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <- deposits:
			balance += amount
		case withdraw := <- withdrawals:
			if withdraw.amount > balance {
				withdraw.isValid <- false
			} else {
				balance -= withdraw.amount
				withdraw.isValid <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
