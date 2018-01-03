package router1

import (
	"fmt"
)

func showCost() {
	fmt.Printf("=========== cost ===========\n")
	for k1, v1 := range cost {
		fmt.Printf("[%d]: ", k1)
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)", k2, v2)
		}
		fmt.Printf("\n")
	}
	if len(cost) == 0 {
		fmt.Printf("(null)\n")
	}
}

func showDist() {
	fmt.Printf("=========== dist ===========\n")
	for k, v := range dist {
		fmt.Printf("(%d,%d)", k, v)
	}
	if len(dist) == 0 {
		fmt.Printf("(null)")
	}
	fmt.Printf("\n")
}

func showPrev() {
	fmt.Printf("=========== prev ===========\n")
	for k, v := range prev {
		fmt.Printf("(%d,%d)", k, v)
	}
	if len(prev) == 0 {
		fmt.Printf("(null)")
	}
	fmt.Printf("\n")
}

func showInfo() {
	showDist()
	showPrev()
	showCost()
}
