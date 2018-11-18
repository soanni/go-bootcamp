package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Println(pow)
	for i := range pow {
		fmt.Println(i)
		pow[i] = 1 << uint(i)
		if pow[i] >= 16 {
			break
		}
	}
	//for _, value := range pow {
	//	fmt.Printf("%d\n", value)
	//}
	fmt.Println(pow)
}
