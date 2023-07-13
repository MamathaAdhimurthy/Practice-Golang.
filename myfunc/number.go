package myfunc

import (
	"fmt"
	"strconv"
)

func NumberOrNot() {
	x := "1234"
	//x:="123b"
	val, err := strconv.Atoi(x)
	if err != nil {
		fmt.Printf("supplied value %s is not a number\n", x)
	} else {
		fmt.Printf("Suplied value %s is a number with value %d\n", x, val)
	}
}
