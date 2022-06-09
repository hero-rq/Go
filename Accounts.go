package accounts 

import (
	"fmt"
	"errors"
	)

type Account struct {
	owner string
	balance int
}

func NewAccounts(owner string) *Account {
	account := Account{owner:owner, balance:0}
	return &account
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("you are making the new universe")
	}
	a.balance -= amount
	return nil
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \n has: ", a.Balance())
}
