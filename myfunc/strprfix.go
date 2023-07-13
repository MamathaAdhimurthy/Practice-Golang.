package myfunc

import (
	"fmt"
	"strings"
)

func StrPrefix() {
	res := strings.HasPrefix("abcd", "ab")
	fmt.Println(res)
	res = strings.HasPrefix("abcd", "ac")
	fmt.Println(res)
	res1 := strings.TrimSuffix("asdfgh", "gh")
	fmt.Println(res1)

}
