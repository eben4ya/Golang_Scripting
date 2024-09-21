package main

import "fmt"

func main(){
	// Array in Go
	var myArray [5]int
	myArray = [5]int{1, 2, 3, 4, 5}

	fmt.Println(myArray)

	myArray2 := [10]int{1, 2, 3, 4, 5}
	// Error
	// myArray2[11] = 6  
	fmt.Println(myArray2)
}