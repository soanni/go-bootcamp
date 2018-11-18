package main

import "fmt"

type Bootcamp struct {
	Lat float64
	Lon float64
}

func main() {
	x := new(Bootcamp)
	y := &Bootcamp{}
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Println(*x == *y)
}
