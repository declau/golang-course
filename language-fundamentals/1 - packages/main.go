package main

import (
	"fmt"
	"modulo/helper"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("From file main")
	helper.Write()

	erro := checkmail.ValidateFormat("dec@gmail.com")
	fmt.Println(erro)
}
