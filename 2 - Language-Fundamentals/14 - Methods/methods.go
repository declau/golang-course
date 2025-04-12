package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Save data from user %s on database\n", u.name)
}

func (u user) majority() bool {
	return u.age >= 18
}

func (u *user) birthday() {
	u.age++
}

func main() {

	user1 := user{"User 1", 33}
	fmt.Println(user1)
	user1.save()

	user2 := user{"Dec", 17}
	user2.save()

	majority := user2.majority()
	fmt.Println(majority)

	user2.birthday()
	fmt.Println(user2.age)

}
