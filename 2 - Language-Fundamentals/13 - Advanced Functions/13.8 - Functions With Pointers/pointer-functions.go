package main

import "fmt"

// Without Pointer
func signalInverter(number int) int {
	return number * -1
}

// With Pointer
func signalInverterWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	fmt.Println("Pointer Functions!")

	number := 20
	inverterNumber := signalInverter(number)
	fmt.Println(inverterNumber)
	fmt.Println(number)

	newNumber := 40
	fmt.Println(newNumber)
	signalInverterWithPointer(&newNumber)
	fmt.Println(newNumber)
}
