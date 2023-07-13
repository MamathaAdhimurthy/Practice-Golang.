package myfunc

import (
	"fmt"
	"strings"
)

func StringReplace() {
	// res := strings.Replace("abcdabxyabr", "ab", "12", 1)
	// fmt.Println(res)
	// res = strings.Replace("abcdabxyabr", "ab", "12", -1)
	// fmt.Println(res)
	res1 := strings.LastIndex("abcdef", "bc")
	fmt.Println(res1)

	//Output will be 8 as "cd" is present in "abcdefabcdef" at index 8
	res1 = strings.LastIndex("abcdefabcdef", "cd")
	fmt.Println(res1)

	//Output will be -1 as "ba" is not present in "abcdef"
	res1 = strings.LastIndex("abcdef", "ba")
	fmt.Println(res1)

}
