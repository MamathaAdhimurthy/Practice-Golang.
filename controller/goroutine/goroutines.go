package main

import myfunc "example.com/hello/myfun"

//	switch FOR COntroller
//
// Select() for channel
// func Seleect(c chan int, quits chan struct{}) {
// 	for {
// 		time.Sleep(time.Second)
// 		select {
// 		case <-c:
// 			fmt.Println("received")
// 		case <-quits:
// 			fmt.Println("quit")
// 			os.Exit(0)
// 		}
// 	}
// }

func main() {
	// wait groups
	//wg := &sync.WaitGroup{}
	// // add,done,& wait
	// wg.Add(2)
	// go func() {
	// 	fmt.Println("Hello")
	// 	wg.Done()
	// }()
	// go func() {
	// 	fmt.Println("World")
	// 	wg.Done()
	// }()
	// fmt.Println("fin")
	// wg.Wait()
	// fmt.Println("Exit")
	// //select {} //it will run infinetly

	//SYNTAX FOR Channels
	//unBUFFERED CHANNELS
	// c := make(chan int)
	// //send  in a go routine
	// go func() {
	// 	c <- 1
	// }()

	// //sinff
	// val := <-c
	// fmt.Println(val)

	// //send
	// go func() {
	// 	c <- 2
	// }()
	// //sniff
	// time.Sleep(time.Second * 2)
	// val = <-c
	// fmt.Println(val)
	// fmt.Println(c)

	//BUFFERED CHANNELS

	// 	type Car struct {
	// 		Model string
	// 	}
	// 	c := make(chan *Car, 3)

	// 	go func() {
	// 		c <- &Car{"1"}
	// 		c <- &Car{"2"}
	// 		c <- &Car{"3"}
	// 		c <- &Car{"4"}
	// 		close(c)
	// 	}()
	// 	for i := range c {
	// 		fmt.Println(i.Model)
	// 	}

	//SELECT BLOCK

	// c := make(chan int, 2)
	// quits := make(chan struct{})

	// go func() {
	// 	c <- 1
	// }()
	// go Seleect(c, quits)
	// select {}

	//USING WG AND CHANNEL IN OTHER BLOCKS PASSING AS PARAMETERS

	// 	}
	myfunc.Even()
}
