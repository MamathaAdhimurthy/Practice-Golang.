package myfunc

import (
	"fmt"
)

type Cars interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Drive() {
	fmt.Println("lambo is moving")
	fmt.Println(l.LamboModel)
}
func (c *Chevy) Drive() {
	fmt.Println("Chevy is moving")
	fmt.Println(c.ChevyModel)
}

func main() {
	l := Lambo{"Gallardo"}
	c := Chevy{"c345"}
	l.Drive()
	c.Drive()
}
