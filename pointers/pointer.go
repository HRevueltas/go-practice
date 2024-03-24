package pointers

import "fmt"

// PointerExample is a function that demonstrates the use of pointers in Go
func PointerExample() {
	// Declare a variable of type int
	var number int = 10

	// Declare a pointer variable of type int
	var pointer *int

	// Assign the address of number to the pointer variable
	pointer = &number

	// Access the value of the variable using the pointer
	fmt.Println("Value of the variable:", number)
	fmt.Println("Address of the variable:", &number)
	fmt.Println("Value of the variable using the pointer:", *pointer)
	fmt.Println("Address of the variable using the pointer:", pointer)
}
