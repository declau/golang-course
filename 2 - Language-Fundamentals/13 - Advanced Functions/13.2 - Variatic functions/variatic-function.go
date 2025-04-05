package main

import "fmt"

func add(numbers ...int) int {
	fmt.Println(numbers)
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total

}

func main() {
	totalAdd := add(1, 2, 3)
	fmt.Println(totalAdd)
}
