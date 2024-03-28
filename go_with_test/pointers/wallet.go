package pointers

import "fmt"

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(balance Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += balance
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
