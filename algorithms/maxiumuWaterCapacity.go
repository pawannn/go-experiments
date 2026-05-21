package main

func MaxWaterStore(inp []int) int {
	max := 0

	ptr1 := 0
	ptr2 := len(inp) - 1

	for ptr1 < ptr2 {
		minHeight := min(inp[ptr1], inp[ptr2])

		distance := ptr2 - ptr1

		capacity := minHeight * distance
		if capacity > max {
			max = capacity
		}

		if minHeight == inp[ptr1] {
			ptr1++
		} else {
			ptr2--
		}
	}

	return max
}
