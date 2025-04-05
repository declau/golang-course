package main

import "fmt"

func function1() {
	fmt.Println("Exec function 1!")

}

func function2() {
	fmt.Println("Exec function 2!")
}

func studentApproved(n1, n2 float32) bool {
	defer fmt.Println("result will return")
	fmt.Println("Enter in the function to check if student is appoved!")
	media := (n1 + n2) / 2

	if media >= 6 {

		return true
	}

	return false

}

func main() {

	//defer function1()
	//function2()

	fmt.Println(studentApproved(7, 8))

}
