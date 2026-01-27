package main

import (
	"fmt"
	"log"
)

func main() {
	token, err := GenerateToken("12345", "pawan@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(token)

	claims, err := ParseToken(token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", *claims)
}
