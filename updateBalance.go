package main

import (
	"errors"
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(c *customer, tr transaction) error {
	if tr.transactionType == transactionDeposit {
		c.balance += tr.amount
	} else if tr.transactionType == transactionWithdrawal {
		if c.balance < tr.amount {
			return errors.New("insufficient funds")
		}
		c.balance -= tr.amount
	} else {
		return errors.New("unknown transaction type")
	}
	return nil
}

func main() {
	alice := customer{id: 1, balance: 100.0}
	deposit := transaction{customerID: 1, amount: 50, transactionType: transactionWithdrawal}

	fmt.Println(updateBalance(&alice, deposit))
	fmt.Println(alice)
}
