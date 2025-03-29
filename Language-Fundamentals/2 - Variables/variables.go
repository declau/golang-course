package main

import "fmt"

func main() {
	var var1 string = "variable 1"
	fmt.Println(var1)

	var2 := "variable 2"
	fmt.Println(var2)

	var (
		var3 string = "varialbe 3"
		var4 string = "variable 4"
	)
	fmt.Println(var3, var4)

	var5, var6 := "variable 5", "varialbe 6"
	fmt.Println(var5, var6)

	const constant1 string = "constant 1"
	fmt.Println(constant1)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)

}
