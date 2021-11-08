package main

import (
	"fmt"
	"strconv"
)

func is_palimdrome(n int64) bool {
	s := strconv.FormatInt(n, 10)

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {
	largest := int64(0)
	i, j := int64(999), int64(999)

	for i > 0 && j > 0 {
		product := i * j

		if is_palimdrome(product) && product > largest {
			largest = product
			i--
			j = 99
			if i*j < largest {
				break
			}
		} else if j > 1 {
			j -= 1
		} else {
			i--
			j = 999
		}

	}
	fmt.Println(largest)
}
