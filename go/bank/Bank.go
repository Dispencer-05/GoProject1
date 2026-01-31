package bank

import (
	"strings"

	"github.com/Dispencer-05/GoProject1/tree/main/go/account"
)

type Bank struct {
	accounts []account.Account
}

func NewBank() *Bank {
	return &Bank{
		accounts: make([]account.Account, 0),
	}
}

func (b *Bank) Add(a account.Account) {
	b.accounts = append(b.accounts, a)
}

func (b *Bank) Accrue(rate float64) {
	for _, a := range b.accounts {
		a.Accrue(rate)
	}
}

func (b *Bank) String() string {
	var sb strings.Builder
	for _, a := range b.accounts {
		sb.WriteString(a.String())
		sb.WriteByte('\n')
	}
	return sb.String()
}
