package myfunc

import (
	"fmt"
	"strings"
)

func SplitString() {
	res := strings.Split("ab$cd$ef", "$")
	fmt.Println(res)
	res1 := strings.Split("ab$cd$ef", "-")
	fmt.Println(res1)
	res2 := strings.Split("ab$cd$ef", "")
	fmt.Println(res2)
	res3 := strings.Split("", "")
	fmt.Println(res3)

}
