package main

import "fmt"

func main() {
	a := [2]string{"hello", "world!"}
	fmt.Println(a)
	fmt.Printf("%s\n", a)
	fmt.Printf("%q\n", a)
}
