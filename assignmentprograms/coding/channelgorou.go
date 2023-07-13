package coding

import (
	"fmt"
)

func Send(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}

}

func Receive(ch <-chan int) {
	fmt.Println("Printing numbers from 1 to 10:")
	for num := range ch {
		fmt.Println(num)
	}
}
