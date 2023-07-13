package channelassignment

import (
	"fmt"
)

func Send2(ch chan int) {
	fmt.Println("Sending value to channnel start")
	ch <- 1
	fmt.Println("Sending value to channnel finish")

}
