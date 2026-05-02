package main

import (
	"errors"
	"fmt"
)

// Encapsulation basically means hiding internal details of an object and only exposing what is required.
//
// The account number and balance will not be able to access when the struct bank account is used in another package.

type BankAccount struct {
	accountNumber int64   // unexported
	balance       float64 // unexported
}

func NewBankAccount(accNum int64, balance float64) *BankAccount {
	var initialBalance float64 = 0
	if balance > 0 {
		initialBalance = balance
	}

	return &BankAccount{
		accountNumber: accNum,
		balance:       initialBalance,
	}
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return errors.New("0 amount to withdraw")
	}

	if b.balance < amount {
		return errors.New("low balance")
	}

	b.balance -= amount
	return nil
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

func (b *BankAccount) AccountNumber() int64 {
	return b.accountNumber
}

func EncapsulationBankAccount() {
	b := NewBankAccount(10234527890, 200)

	fmt.Println("Account number : ", b.AccountNumber())
	fmt.Println("Balance : ", b.Balance())

	b.Withdraw(100)
	fmt.Println("Balance : ", b.Balance())
}
