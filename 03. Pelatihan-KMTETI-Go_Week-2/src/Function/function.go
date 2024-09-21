package main

import (
	"fmt"
)

// Function in Go
func sayHello(){
	fmt.Println("Hello World")
}

func addNumbers(a float32, b float32) float32{
	return a + b
}

func mutliplyNumbers(a float32, b float32) float32{
	res := a * b
	return res
}

func divideNumbers(a int, b int) (int, int){
	div := a / b
	rem := a % b
	return div, rem
}

func main(){
	// sayHello()

	// func addNumbers
	a := 5.0
	b := 10.0
	fmt.Println(addNumbers(float32 (a),float32(b)))

	// func divideNumbers
	mult := mutliplyNumbers(float32(a),float32(b))
	fmt.Println(mult)

	// func divideNumbers
	div, rem := divideNumbers(10, 3)
	fmt.Println(div, rem)

}