package main

import (
	"fmt"
)

func main() {
	// Arithmetic
	addtion := 1 + 1
	subtraction := 2 - 1
	multiplication := 2 * 2
	division := 10 / 2
	restDivision := 10 % 2

	fmt.Println(addtion, subtraction, multiplication, division, restDivision)

	var num1 int16 = 10
	var num2 int32 = 20
	add := num1 + int16(num2)
	fmt.Println(add)

	// Attribution
	var var1 string = "String"
	var2 := "String2"
	fmt.Println(var1, var2)

	// relational
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// Logical
	fmt.Println("-------------------------------------------------")
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
