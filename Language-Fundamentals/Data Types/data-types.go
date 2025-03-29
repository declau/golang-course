package main

import (
	"fmt"
)

func main() {
	//
	var num int16 = 100
	fmt.Println(num)

	num2 := 20000
	fmt.Println(num2)

	var num3 uint32 = 30000
	fmt.Println(num3)

	// Alias
	// int32 == RUNE
	var num4 rune = 33000
	fmt.Println(num4)
	// int8 ==  BYTE
	var num5 byte = 123
	fmt.Println(num5)

	var numReal float32 = 12.345
	fmt.Println(numReal)

	var numReal2 float64 = 123.456
	fmt.Println(numReal2)

	numReal3 := 1234.5678
	fmt.Println((numReal3))

	var str string = "text"
	fmt.Println(str)

	str2 := "text2"
	fmt.Println(str2)

	// There is no character char
	char := 'B'
	fmt.Println(char)

	// Boolean initial value in Golang is false
	var boolean1 bool = false
	fmt.Println(boolean1)

	var error error
	fmt.Println(error)

}
