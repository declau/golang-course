package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	number := 10
	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Less than 15")
	}

	if number2 := number; number2 > 0 {
		fmt.Println("Greater than 0")
	} else if number2 < -10 {
		fmt.Println("Less than -10")
	} else {
		fmt.Println("Between 0 and -10")
	}

}
