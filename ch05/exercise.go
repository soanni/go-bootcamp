package main

import (
	"fmt"
	"strings"
)

func main() {
	var numberOfCoinsDistributed int
	totalCoins := 50
	users := [10]string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution := make(map[string]int, len(users))

	distribute := func(name string) int {
		characters := strings.Split(strings.ToLower(name), "")
		// fmt.Println(characters)
		coins := 0
		for _, v := range characters {
			switch v {
			case "a":
				coins += 1
				break
			case "e":
				coins += 1
				break
			case "i":
				coins += 2
				break
			case "o":
				coins += 3
				break
			case "u":
				coins += 4
			}
		}

		return coins
	}

	for _, v := range users {
		numberOfCoinsDistributed = distribute(v)
		if numberOfCoinsDistributed > 10 {
			numberOfCoinsDistributed = 10
		}
		distribution[v] = numberOfCoinsDistributed
		totalCoins -= numberOfCoinsDistributed
	}
	fmt.Println(distribution)
	fmt.Printf("Coins left: %d", totalCoins)
}
