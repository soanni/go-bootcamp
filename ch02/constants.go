package main

import "fmt"

const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 10
	Small = Big >> 9
)

func main() {
	const Greeting = "Hello world"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
	fmt.Println(Small)
}
