package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

// Generics let you write functions or types that work with multiple data types while preserving type safety.
func sumNumbers[T Number](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func AddSlices() {
	sum1 := sumNumbers([]int{1, 2, 3, 4})
	sum2 := sumNumbers([]float64{1.1, 2.1, 3.1, 4.1})

	fmt.Println(sum1)
	fmt.Println(sum2)
}
