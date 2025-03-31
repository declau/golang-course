package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var1 := 10
	var2 := var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)
}
