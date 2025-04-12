package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Save data from user %s on database", u.name)
}

func main() {

	user1 := user{"User 1", 33}
	fmt.Println(user1)
	user1.save()

}
