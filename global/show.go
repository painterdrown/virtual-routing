package global

import (
	"fmt"
)

// ShowCost .
func ShowCost() {
	for k1, v1 := range Cost {
		fmt.Printf("[%d]: ", k1)
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)", k2, v2)
		}
		fmt.Printf("\n")
	}
}

// ShowDist .
func ShowDist() {
	fmt.Printf("[%d]: ", Port)
	for k, v := range Dist {
		fmt.Printf("(%d,%d)", k, v)
	}
	fmt.Printf("\n")
}
