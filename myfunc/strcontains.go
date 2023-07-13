package myfunc

import (
	"fmt"
	"strings"
)

func StringContainsOrNot() {
	res := strings.Contains("abc", "bc")
	fmt.Println(res)
	res1 := strings.Contains("abc", "de")
	fmt.Println(res1)
}
