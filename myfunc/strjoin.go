package myfunc

import (
	"fmt"
	"strings"
)

func StrJoin() {
	res := strings.Join([]string{"ab", "cd", "ef"}, "-")
	fmt.Println(res)
	res = strings.Join([]string{}, "-")
	fmt.Println(res)
	res = strings.Join([]string{"ab", "cd", "ef"}, "")
	fmt.Println(res)
}
