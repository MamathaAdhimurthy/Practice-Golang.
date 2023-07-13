package myfunc

import (
	"fmt"
	"strings"
)

func StringLower() {
	res := strings.ToLower("ABC")
	fmt.Println(res)
	res = strings.ToLower("ABC12$a")
	fmt.Println(res)
	res = strings.ToUpper("abc")
	fmt.Println(res)
	//This is sentence TITLE
	res = strings.Title("this is a test sentence")
	fmt.Println(res)

}
