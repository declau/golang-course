package main

import "fmt"

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

type form interface {
	area() float64
}

func main() {
	fmt.Println("Interfaces")
}
