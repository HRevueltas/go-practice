// package variables
package main

import "fmt"

func main() {
	variables()
}

var r = 5

func variables() {
	fmt.Println(r)
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
