package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	sortMap := make(map[int][]string)
	for _, v := range names {
		sortMap[len(v)] = append(sortMap[len(v)], v)
	}

	// fmt.Printf("%+v\n", sortMap)
	for i, v := range sortMap {
		fmt.Printf("Slice of names with length %d: %q\n", i, v)
	}
}
