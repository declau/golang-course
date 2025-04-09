package main

import "fmt"

func signalInverter(number int) int {
	return number * -1
}

func main() {
	fmt.Println("Pointer Functions!")
	number := 20
	inverterNumber := signalInverter(number)
	fmt.Println(inverterNumber)
	fmt.Println(number)
}
