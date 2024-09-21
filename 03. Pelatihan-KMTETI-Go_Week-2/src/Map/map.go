package main

import "fmt"

func main(){
	// Map in Go
	myMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
	}
	// print string in alphabetical order
	fmt.Println(myMap)
	// print value of key "one"
	fmt.Println(myMap["one"])

	myMap2 := map[int]int{
		3: -4,
		4: -3,
		2: -2,
	}
	// print int in ascending order
	fmt.Println(myMap2)
}