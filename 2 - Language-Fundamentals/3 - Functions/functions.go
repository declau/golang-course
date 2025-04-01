package main

import "fmt"

func add(n1 int8, n2 int8) int8 {
	return n1 + n2

}

func mathematicalCalculations(n1, n2 int8) (int8, int8) {
	add := n1 + n2
	subtraction := n1 - n2
	return add, subtraction

}

func main() {
	adding := add(10, 20)
	fmt.Println(adding)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Text fron function f")

	addResult, subtractionresult := mathematicalCalculations(10, 20)
	fmt.Println(addResult, subtractionresult)

	addResult2, _ := mathematicalCalculations(10, 20)
	fmt.Println(addResult2)
}
