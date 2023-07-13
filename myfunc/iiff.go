package myfunc

import (
	"fmt"
)

func Iiff() {
	squareOf2 := func() int {
		return 2 * 2
	}()
	fmt.Println(squareOf2)
}
