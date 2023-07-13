package channelassignment

import (
	"fmt"
	"time"
)

// To illustrate blocking on send lets add a log after we send the value to channel ch in the send() function and a timeout in the receive() function before we receive value from ch
// func Send1(ch chan int) {
// 	ch <- 1
// 	fmt.Println("Sending value to the channel ")

// }

// func Receive1(ch chan int) {
// 	//time.Sleep(time.Second * 2)
// 	fmt.Println("Receiving value from the channel")
// 	_ = <-ch
// 	return

// }

// To illustrate blocking on receive lets add a log after we receive the value in the recieve() function and a timeout in the send() function before we send value
func Send1(ch chan int) {
	//time.Sleep(time.Second * 1)
	fmt.Println("Timeout finished")
	ch <- 1
}

func Receive1(ch chan int) {
	time.Sleep(time.Second * 1)
	val := <-ch
	fmt.Println("Receiving value from the channel", val)

}
