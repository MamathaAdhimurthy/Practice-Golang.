package myfunc

import (
	"fmt"
	"strings"
)

func CompareString() {
	res := strings.Compare("abc", "abc")
	fmt.Println(res)
	res1 := strings.Compare("abc", "xyz")
	fmt.Println(res1)
	res2 := strings.Compare("xyz", "abc")
	fmt.Println(res2)

}
