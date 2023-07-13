package myfunc

import (
	"fmt"
)

// passing function as an argument to another function
func Print(f func(int, int) int, a, b int) {
	fmt.Println(f(a, b))
}
func Area(a, b int) int {
	return a * b
}
func Sum1(a, b int) int {
	return a + b
}

//Return a function from function

// func GetAreaFunc() func(int, int) int {
// 	return func(x, y int) int {
// 		return x * y
// 	}
// }

// PASS VARIABLE NUMBER OF ARGUMENTS TO A FUNCTION
// Different number of parameters but of the same type
func Add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

//Different number of parameters and of different types

func Handle(params ...interface{}) {
	fmt.Println("Handle function called with parameters:")
	for _, param := range params {
		fmt.Printf("%v\n", param)
	}

}
