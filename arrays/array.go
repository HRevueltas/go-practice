package arrays

import (
	"fmt"
	"time"
)

func Array() {
	start := time.Now()

	var a [5]int
	// this is an array of 5 integers
	fmt.Println("emp:", a)

	// set a value at an index
	a[2] = 100
	fmt.Println("set:", a)
	// get a value at an index
	fmt.Println("get:", a[2])
	// get the length of an array
	fmt.Println("len:", len(a))

	// declare and initialize an array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// two-dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("Duration: ", duration)

}

func MostHigh() {
	// find the highest value in an array
	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println("numbers: %d", numbers)
}
