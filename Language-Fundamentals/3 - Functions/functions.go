package main

import "fmt"

func add(n1 int8, n2 int8) int8 {
	return n1 + n2

}

func main() {
	adding := add(10, 20)
	fmt.Println(adding)
}
