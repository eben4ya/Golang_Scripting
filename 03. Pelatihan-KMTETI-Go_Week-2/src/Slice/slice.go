package main

import "fmt"

func main(){
	// Slice in Go
	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Println(mySlice)

	mySlice = append(mySlice, 6)
	fmt.Println(mySlice)
	
}