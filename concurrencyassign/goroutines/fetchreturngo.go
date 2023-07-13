package channelassign

func Sum(a, b int, result chan int) {
	sumValue := a + b
	result <- sumValue //push sum value to result
	return

}
