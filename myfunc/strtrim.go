package myfunc

import (
	"fmt"
	"strings"
)

func StringTrim() {
	res := strings.TrimSpace(" test ")
	fmt.Println(res)
	res = strings.TrimSpace(" This is test ")
	fmt.Println(res)

}
