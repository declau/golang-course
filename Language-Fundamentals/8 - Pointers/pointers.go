package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var1 := 10
	var2 := var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	// Pointers
	var var3 int
	var pointer *int

	var3 = 100
	pointer = &var3
	fmt.Println(var3, *pointer) //The * is a dereferencing process

	var3 = 150
	fmt.Println(var3, *pointer) //The * is a dereferencing process
}
