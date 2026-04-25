package main

import "fmt"

func SomeSafeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic : ", r)
		}
	}()

	fmt.Println("We are doing something")
	panic("I did something wrong")
	fmt.Println("further ops will not be executed")
}

func startSafFunc() {
	SomeSafeFunction()
	fmt.Println("REMAINING EXECUTION WILL CONTINUE")
}
