package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount create account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on yout account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance
func (a *Account) Balance() int {
	return a.balance
}

var errNomoney = errors.New("Can't withdraw")

// withdraw x
func (a *Account) WithDraw(amount int) error {

	if a.balance < amount {
		return errNomoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account. \nHas: ", a.Balance())
}
