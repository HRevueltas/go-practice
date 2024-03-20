package main

import "fmt"

// import "os/user"

// import "os/user"

// import fmtpackage "project-layout/fmtPackage"

func main() {

	fmt.Println(computePrice(10, 20.5))
}

func computePrice(numberOfItems int, pricePerItem float64) (float64, string) {
	if numberOfItems < 0 {
		return 0, "Invalid number of items"
	}

	return float64(numberOfItems) * pricePerItem, ""

}
