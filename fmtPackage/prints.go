package fmtpackage

import (
	"fmt"
)

func Prints() {
	pi := 3.14159265359
	const name, age = "Kim", 22
	fmt.Print("Text without new line below\n")
	fmt.Println("Text with new line below")
	fmt.Printf("Text with format %v\n", pi)

	fmt.Printf("My name is %v and I am %v years old.\n", name, age)
}
