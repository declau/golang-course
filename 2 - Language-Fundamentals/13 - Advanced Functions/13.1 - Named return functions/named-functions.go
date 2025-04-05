package main

import "fmt"

func mathematicalCalculations(n1, n2 int) (add, sub int) {
	add = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	add, sub := mathematicalCalculations(10, 20)
	fmt.Println(add, sub)
}
