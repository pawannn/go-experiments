package main

import "fmt"

// A closure is a function that captures and remembers variables from its surrounding scope even after that outer scope has finished executing.
func activateGiftCard() func(int) int {
	amount := 100

	debitAmount := func(debitAmount int) int {
		return debitAmount - amount
	}

	return debitAmount
}

func checkout() {
	amount := 200
	giftCard := activateGiftCard()
	totalAoumt := giftCard(amount)
	fmt.Println(totalAoumt)
}
