package main

import (
	"errors"
	"fmt"
)

// In golang, we have composition in place of inheritance.
// Composition in a way by which we build complext types by embeding small structs into the main struct
// By embeding the a struct into another struct, the behaviour of the smaller struct can be resued in the main struct.

type Account struct {
	accountNumber int64
	balance       float64
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) AccountNumber() int64 {
	return a.accountNumber
}

type SavingsAccount struct {
	interestRate float64
	Account
}

func (s *SavingsAccount) AddInterest() {
	s.balance += s.balance * s.interestRate
}

type CurrentAccount struct {
	Account
	overDraftLimit float64
}

func (c *CurrentAccount) Withdraw(amount float64) error {
	if amount > c.balance+c.overDraftLimit {
		return errors.New("Insufficient balance")
	}

	fmt.Println("Amount withdrawn")
	c.balance -= amount
	return nil
}
