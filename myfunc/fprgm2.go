package myfunc

import (
	"fmt"
)

func GetModulus() func(int) int {
	count := 0
	return func(x int) int {

		count = count + 1
		fmt.Printf("modulus function called %d times\n", count)
		if x < 0 {
			x = x * -1
		}
		return x

	}
}
