package coding

import (
	"fmt"
)

func FrequncyOfEle() {
	arr := []int{1, 2, 3, 2, 1, 3, 4, 5, 4}
	//arr1 := []string{"sri", "mamatha", "sri", "dhanuja", "sri", "mamatha"}
	//res := []int{}

	frequencyMap := make(map[int]int)
	//COUNT FREQUENCY OF EACH ELEMENT
	for _, num := range arr {
		frequencyMap[num]++
	}

	//PRINT THE FREQUENCY OF EACH ELEMENT
	for num1, frequency := range frequencyMap {
		fmt.Printf("%d:%d\n", num1, frequency)
		//res = append(res, num)
	}
	//fmt.Println(sort.IntSlice(res))

}
