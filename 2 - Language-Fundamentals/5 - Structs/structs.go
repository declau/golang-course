package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint32
}

func main() {
	fmt.Println("Structs file")

	var u user
	u.name = "Dec"
	u.age = 42
	fmt.Println(u)

	address := address{"Av Test", 0}
	u2 := user{"Dec", 42, address}
	fmt.Println(u2)

	u3 := user{age: 42}
	fmt.Println(u3)
}
