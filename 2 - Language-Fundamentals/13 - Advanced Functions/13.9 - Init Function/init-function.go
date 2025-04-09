package main

import "fmt"

var n int

func init() {
	fmt.Println("Init function being executed")
	n = 10
}

func main() {
	fmt.Println("Main function being executed")
	fmt.Println(n)
}
