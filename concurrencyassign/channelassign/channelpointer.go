package channelassignment

import (
	"fmt"
)

func Process(ch *chan int) {
	*ch <- 2
	//*ch <- 3
	s := <-*ch

	//*ch = nil
	fmt.Println(s)

}
