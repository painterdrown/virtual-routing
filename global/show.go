package global

import (
	"fmt"
)

// ShowCost .
func ShowCost() {
	for k1, v1 := range Cost {
		fmt.Printf("[Cost][%d]: ", k1)
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)", k2, v2)
		}
		fmt.Printf("\n")
	}
}

// ShowDist .
func ShowDist() {
	fmt.Printf("[Dist][%d]: ", Port)
	for k, v := range Dist {
		fmt.Printf("(%d,%d)", k, v)
	}
	fmt.Printf("\n")
}

// ShowPrev .
func ShowPrev() {
	fmt.Printf("[Prev][%d]: ", Port)
	for k, v := range Prev {
		fmt.Printf("(%d,%d)", k, v)
	}
	fmt.Printf("\n")
}

// ShowNear .
func ShowNear() {
	fmt.Printf("[Near][%d]: ", Port)
	for k := range Near {
		fmt.Printf("%d ", k)
	}
	fmt.Printf("\n")
}
