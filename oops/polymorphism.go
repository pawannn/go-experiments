package main

import "fmt"

// Polymorphism means many forms - means, a single same function can perform in different way according to the given context.
// The IMakePayment function is a same function but can perform card payment, upi payment
func IMakePayment(method PaymentMethod, amount float64) error {
	return method.Pay(amount)

}

type PaymentMethod interface {
	Pay(amount float64) error
}

type ICardPayment struct {
	cardNumber int64
}

func (c ICardPayment) Pay(amount float64) error {
	fmt.Println("Paid using card : ", amount)
	return nil
}

type IUpiPayment struct {
	upiID string
}

func (u IUpiPayment) Pay(amount float64) error {
	fmt.Println("Paid using UPI : ", amount)
	return nil
}
