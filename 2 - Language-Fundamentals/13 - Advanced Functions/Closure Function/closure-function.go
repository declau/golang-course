package main

import "fmt"

func closure() func() {
	text := "Inside the closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Inside the main function"
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}
