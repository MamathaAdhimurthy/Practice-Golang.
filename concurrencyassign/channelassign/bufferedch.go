//For range loop on a channel

package channelassignment

import "fmt"

func Sum1(ch chan int) {
	sum := 0
	for val := range ch {
		sum += val

	}
	fmt.Printf("Value of sum:%d\n", sum)
}
