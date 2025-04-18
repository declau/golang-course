package main

import (
	"fmt"
	"reflect"
)

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

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 10, 11, 12, 13}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 14)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Alter Position"
	fmt.Println(array2)
	fmt.Println("---------------------------------------")

	// Internal Arrays

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacity

}
