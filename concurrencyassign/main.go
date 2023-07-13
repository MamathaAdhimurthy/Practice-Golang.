package main

import (
	"fmt"
	"time"

	"example.com/hello/myfun/concurrencyassign/channelassignment"
)

// fetch multiple return values from a goroutine in Go
// type result struct {
// 	SumValue      int
// 	MultiplyValue int
// }

// program to demonstrate that goroutines don’t have parents or children.
func main() {
	// go start() //run in back ground
	// fmt.Println("Started")
	// time.Sleep(1 * time.Second)
	// fmt.Println("Finished")

	//Creating Multiple Goroutines
	// 	for i := 0; i < 10; i++ {
	// 		go execute(i)
	// 	}
	// 	time.Sleep(time.Second * 2)
	// 	fmt.Println("finished")
	// }

	// Creating Multiple Goroutines
	//go channelassign.Start()
	// func execute(id int) {
	// 	fmt.Printf("id : %d\n", id)

	// }

	//fmt.Println(runtime.NumCPU())

	//Get number of currently running/active goroutines

	// for i := 0; i < 20; i++ {
	// 	go channelassign.Execute()

	// }
	// fmt.Println(runtime.NumGoroutine())

	//Wait for all Go routines to finish execution in Golang

	// var wg sync.WaitGroup
	// wg.Add(2)

	// go channelassign.Sleep(&wg, time.Second*1)
	// go channelassign.Sleep(&wg, time.Second*2)
	// wg.Wait()
	// fmt.Println("All go routines finished execution")

	//Fetch Return Value from a goroutine in Go
	// result := make(chan int, 1)        //created variable of result of chnnel length is 1
	// go channelassign.Sum(2, 3, result) //passing channel to sum func
	// value := <-result                  //waiting to collect return value from sum func
	// fmt.Println(value)

	//fetch multiple return values from a goroutine in Go
	// 	resultChan := make(chan result, 1)
	// 	go SumAndMultiply(2, 3, resultChan)
	// 	res := <-resultChan
	// 	fmt.Printf("Sum Value: %d\n", res.SumValue)
	// 	fmt.Printf("Multiply Value: %d\n", res.MultiplyValue)
	// 	close(resultChan)
	// }

	// fetch multiple return values from a goroutine in Go
	// func SumAndMultiply(a, b int, resultChan chan result) {
	// 	SumValue := a + b
	// 	MultiplyValue := a * b
	// 	res := result{SumValue: SumValue, MultiplyValue: MultiplyValue}
	// 	resultChan <- res
	// 	time.Sleep(time.Second * 10)
	// 	return

	//Declare channel

	// var a chan int
	// fmt.Println(a)

	//example of where we will send data from one goroutine and receive that data in another goroutine.

	// ch := make(chan int)
	// fmt.Println("Sending value to the Channel")
	// go channelassignment.Send(ch)
	// fmt.Println("Receving value from channel")
	// go channelassignment.Receive(ch)
	// time.Sleep(time.Second * 1)

	//To illustrate blocking on send lets add a log after we send the value to channel ch in the send() function and a timeout in the receive() function before we receive value from ch
	// ch := make(chan int)
	// go channelassignment.Send1(ch)
	// go channelassignment.Receive1(ch)
	// time.Sleep(time.Second * 2)

	//BUFFERED CHANNEL
	//Send on a channel is blocked when the channel is full
	// ch := make(chan int, 1)
	// ch <- 1
	// //ch <- 1 //Send on a channel is blocked when the channel is full
	// fmt.Println("sending value to channel")
	// val := <-ch
	// //val = <-ch //Receive on a channel is blocked when the channel is empty
	// fmt.Println("receving value fromchannel ", val)

	///For range loop on a channel

	// ch := make(chan int, 3)
	// ch <- 2
	// ch <- 2
	// ch <- 2
	// close(ch)
	// go channelassignment.Sum1(ch)
	// time.Sleep(time.Second * 2)

	// //Nil channel

	// var a chan int
	// fmt.Println(a)

	//channel pointer

	// ch := make(chan int, 3)
	// channelassignment.Process(&ch)
	//fmt.Println(ch)
	//NIL CHANNEL
	// var ch chan int
	// go channelassignment.Send2(ch)
	// <-ch
	// time.Sleep(time.Second * 1)

	//Pause Execution of a goroutine until an activity or event is completed in Go

	ch := make(chan bool)
	go channelassignment.Receivechan(ch)
	//time.Sleep(time.Second * 1)
	fmt.Printf("Send Value: %t\n", true)
	ch <- true

	time.Sleep(time.Second * 2)
}
