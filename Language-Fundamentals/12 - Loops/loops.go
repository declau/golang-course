package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("loops")

	i := 0

	for i < 10 {
		fmt.Println("Increment i")
		time.Sleep(time.Second)
		i++

	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Increment j", j)
		time.Sleep(time.Second)
	}

	names := [3]string{"Dec", "Edu", "Clau"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	user := map[string]string{
		"name":    "Dec",
		"surname": "Clau",
	}
	fmt.Println(user)

	for key, value := range user {
		fmt.Println(key, value)
	}

}
