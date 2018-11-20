package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	fmt.Printf("%T\n", now)
	mins := now % 2
	fmt.Printf("%T\n", mins)
	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	}
}
