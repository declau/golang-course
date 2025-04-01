package main

import "fmt"

type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	person
	course     string
	university string
}

func main() {
	fmt.Print("inheritance")

	p1 := person{"Dec", "Clau", 42, 173}
	fmt.Println(p1)

	s1 := student{p1, "enginner", "USP"}
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.surname)
	fmt.Println(s1.height)

}
