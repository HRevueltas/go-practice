package ifelse

import (
	"fmt"
	"time"
)

func Condition() {
	start := time.Now()
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("Duration: ", duration)
}
