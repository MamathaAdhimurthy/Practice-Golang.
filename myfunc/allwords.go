package myfunc

import (
	"fmt"
	"strings"
)

func AllWords() {
	res := strings.Fields("ab cd ef")
	fmt.Println(res)
	res = strings.Fields("abcdef")
	fmt.Println(res)
	res = strings.Fields("ab  cd  ef")
	fmt.Println(res)
	res = strings.Fields(" ")
	fmt.Println(res)
	res = strings.Fields("")
	fmt.Println(res)
}
