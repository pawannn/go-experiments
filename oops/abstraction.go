package main

import (
	"errors"
	"fmt"
)

// Abstraction is the method of showing essential features only and hiding complex implementation

type Payment interface {
	Pay(float64) error
}

type CardPayment struct {
	cardNumber int64
}

func (c CardPayment) Pay(amount float64) error {
	if !validateCard(c.cardNumber) {
		return errors.New("Invalid card")
	}

	if amount < 0 {
		return errors.New("0 amount to withdraw")
	}

	fmt.Println("Paid using card : ", amount)
	return nil
}

type UPIPayment struct {
	upiID string
}

func (u UPIPayment) Pay(amount float64) error {
	if !validateUPI(u.upiID) {
		return errors.New("invalid upi ID")
	}

	if amount < 0 {
		return errors.New("0 amount to withdraw")
	}

	fmt.Println("Paid using UPI : ", amount)
	return nil
}

// Lhun Algo
func validateCard(cardNum int64) bool {
	return true
}

// Regex
func validateUPI(upi string) bool {
	return true
}

func MakePayment(method Payment, amount float64) error {
	return method.Pay(amount)
}

func StartAbstractPayment() {
	cc := CardPayment{cardNumber: 45672182722881}
	up := UPIPayment{upiID: "pawan@hdfc.yes"}

	MakePayment(cc, 20)
	MakePayment(up, 20)
}
