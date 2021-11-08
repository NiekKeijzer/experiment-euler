package main

import "fmt"

func divisible_by_all(n int64, seq []int64) bool {
	for _, i := range seq {
		if (n % i) != 0 {
			return false
		}
	}

	return true
}

func main() {
	// The challenge states 2520 is the lowest possible for numbers 1 to 10
	i := int64(2520)
	seq := []int64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for {
		if divisible_by_all(i, seq) {
			break
		}

		i += 10
	}

	fmt.Println(i)
}
