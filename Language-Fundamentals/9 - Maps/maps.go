package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":    "Dec",
		"surname": "Clau",
	}
	fmt.Println(user)
	fmt.Println(user["name"])

}
