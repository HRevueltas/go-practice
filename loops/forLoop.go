package main

import (
	"fmt"
	"time"
)

func main() {
	// record the start time
	start := time.Now()

	// call the loops function
	loops()

	// record the end time
	end := time.Now()

	// calculate the duration
	duration := end.Sub(start)

	// print the duration
	fmt.Println("Duration: ", duration)

}

func loops() {
	i := 1
	// for loop for a single condition
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for loop for a classic initial/condition/after
	for j := 0; j <= 3; j++ {
		fmt.Println(j)
	}

	// for range loop
	for i := range 3 {
		fmt.Println(i)
	}

	// for loop with break
	for {
		fmt.Println("loop")
		break
	}

	// for loop with continue
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)

	}
}
