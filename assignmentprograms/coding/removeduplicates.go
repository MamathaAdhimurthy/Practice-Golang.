package coding

import (
	"fmt"
	"sort"
)

func RemoveDuplicates() {
	arr1 := []int{1, 2, 3, 2, 4, 5, 4, 6}
	//Declare array to store the uniq values
	result := []int{}
	//COUNT FREQUENCY OF EACH ELEMENT
	frequencyofMap := make(map[int]int)
	for _, num := range arr1 {
		frequencyofMap[num]++
	}
	//PRINT THE FREQUENCY OF EACH ELEMENT
	for num, _ := range frequencyofMap {
		//fmt.Printf("%d", num)

		//collect uniq keys and append to the empty array result
		result = append(result, num)

	}
	fmt.Println("Removed duplicates from map:", sort.IntSlice(result))
}
