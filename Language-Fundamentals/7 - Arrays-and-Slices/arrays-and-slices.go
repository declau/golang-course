package main

import "fmt"

func main() {
	fmt.Println("Arrays and Slices")

	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	array2 := [5]string{"position 1", "position 2", "position 3", "position 4", "position 5"}
	fmt.Println(array2)

}
