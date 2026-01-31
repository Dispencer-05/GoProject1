package account

import (
	"fmt"

	"customer"
)

type Account interface {
	Balance() float64
	Accrue(rate float64)
	Deposit(amount float64)
	Withdraw(amount float64)
	String() string
}

type BaseAccount struct {
	number   int
	customer *customer.Customer
	balance  float64
}

func (a *BaseAccount) Balance() float64 {
	return a.balance
}

func (a *BaseAccount) Deposit(amount float64) {
	a.balance += amount
}

func (a *BaseAccount) Withdraw(amount float64) {
	a.balance -= amount
}

func (a *BaseAccount) String() string {
	return fmt.Sprintf("%d: %s: %f", a.number, a.customer, a.balance)
}
