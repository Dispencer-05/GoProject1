package account

import "github.com/Dispencer-05/GoProject1/tree/main/go/customer"

type SavingsAccount struct {
	BaseAccount
	interest float64
}

func NewSavingsAccount(number int, customer *customer.Customer, balance float64) *SavingsAccount {
	return &SavingsAccount{
		BaseAccount: BaseAccount{
			number:   number,
			customer: customer,
			balance:  balance,
		},
		interest: 0,
	}
}

func (s *SavingsAccount) Accrue(rate float64) {
	earned := s.balance * rate
	s.interest += earned
	s.balance += earned
}
