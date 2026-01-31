package account

import "github.com/Dispencer-05/GoProject1/tree/main/go/customer"

type CheckingAccount struct {
	BaseAccount
}

func NewCheckingAccount(number int, customer *customer.Customer, balance float64) *CheckingAccount {
	return &CheckingAccount{
		BaseAccount: BaseAccount{
			number:   number,
			customer: customer,
			balance:  balance,
		},
	}
}

// Checking accounts don't accrue interest.
func (c *CheckingAccount) Accrue(rate float64) {
	// no-op
}
