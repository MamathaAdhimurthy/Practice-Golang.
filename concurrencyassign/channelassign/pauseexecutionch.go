package channelassignment

import (
	"fmt"
)

func Receivechan(ch chan bool) {
	value := <-ch
	fmt.Printf("Received Value: %t\n", value)
	fmt.Println("Doing some work")
}
