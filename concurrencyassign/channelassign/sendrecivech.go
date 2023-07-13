package channelassignment

import (
	"fmt"
)

func Send(ch chan int) {
	ch <- 1
}
func Receive(ch chan int) {
	val := <-ch
	//fmt.Println("value received = %d in receive function\n", val)
	fmt.Printf("Value Received=%d in receive function\n", val)
}
