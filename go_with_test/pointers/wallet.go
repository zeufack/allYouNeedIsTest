package pointers

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(balance int) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += balance
}

func (w *Wallet) Balance() int {
	return w.balance
}
