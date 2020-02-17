package main

import "fmt"

func main() {
	fmt.Println("start")
	number1 := 1
	number2 := 2

	total := add(number1, number2)

	fmt.Println(total)
	fmt.Println("end")
}

// add .
func add(num1, num2 int) int {
	fmt.Println("addtest")
	return num1 + num2
}
