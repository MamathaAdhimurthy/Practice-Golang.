package channelassign

import (
	"fmt"
)

func Start() {
	go Start2()
	fmt.Println("In Goriuttine")
}

func Start2() {
	panic("In Goroutine2")
}
