package main

import (
	"fmt"

	"github.com/Dispencer-05/GoProject1/tree/main/go/account"
	"github.com/Dispencer-05/GoProject1/tree/main/go/bank"
	"github.com/Dispencer-05/GoProject1/tree/main/go/customer"
)

func main() {
	b := bank.NewBank()

	ann := customer.NewCustomer("Ann")
	bob := customer.NewCustomer("Bob")

	b.Add(account.NewCheckingAccount(1, ann, 100.00))
	b.Add(account.NewSavingsAccount(2, ann, 200.00))
	b.Add(account.NewSavingsAccount(3, bob, 150.00))

	b.Accrue(0.02)

	fmt.Println(b)
}
