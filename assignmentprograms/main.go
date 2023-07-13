package main

import (
	"time"

	"example.com/hello/myfun/assignment/coding"
)

func main() {
	// 	s1 := coding.Student{
	// 		Name:   "Sri",
	// 		Age:    21,
	// 		Grades: []float64{80.50, 90.0, 75.5, 88.0},
	// 	}
	// 	s2 := coding.Student{
	// 		Name:   "Mamatha",
	// 		Age:    23,
	// 		Grades: []float64{92.0, 87.5, 95.5, 89.0},
	// 	}
	// 	fmt.Printf("Average grade for %s:%.2f\n", s1.Name, s1.CalculateAverageGrade())
	// 	fmt.Printf("Average grade for %s:%.2f\n", s2.Name, s2.CalculateAverageGrade())
	// }

	// Program 2
	// rect := coding.Rectangle{Width: 5, Height: 3}
	// circ := coding.Circle{Radius: 4}
	// tri := coding.Traingle{Base: 5, Height: 3, Side1: 3, Side2: 4, Side3: 5}

	// fmt.Println("Area of Reactangle:", rect.Area())
	// fmt.Println("Perimeter of Rectangle:", rect.Perimeter())

	// fmt.Println("Area of Circle:", circ.Area())
	// fmt.Println("Perimeter of Circle:", circ.Perimeter())

	// fmt.Println("Area of traingle:", tri.Area())
	// fmt.Println("Perimeter of Traingle:", tri.Perimeter())

	//PROGRAM 3

	//coding.FrequncyOfEle()
	//coding.RemoveDuplicates()

	//Task 4
	//PRINT NUMBERS FROM 1 TO 10

	ch := make(chan int)
	go coding.Send(ch)
	go coding.Receive(ch)
	time.Sleep(time.Second * 2)

}
