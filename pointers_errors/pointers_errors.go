package pointers_errors

import (
	"errors"
	"fmt"
)

var ErrUnsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var ErrNotPositiveAmount = errors.New("cannot deposite negative or zero amount")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount <= 0 {
		return ErrNotPositiveAmount
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrUnsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
