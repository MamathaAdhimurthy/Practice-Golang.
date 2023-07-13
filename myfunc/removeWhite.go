package myfunc

import (
	"fmt"
	"strings"
)

func RemoveWhiteSpace() {
	sample := " This is a sample string "
	noSpaceString := strings.ReplaceAll(sample, " ", "")
	fmt.Println(noSpaceString)
}
