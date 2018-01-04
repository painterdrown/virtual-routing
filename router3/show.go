package router3

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

func showNear() {
	fmt.Printf("=========== near ===========\n")
	for k := range near {
		fmt.Printf("%d ", k)
	}
	if len(near) == 0 {
		fmt.Printf("(null)")
	}
	fmt.Printf("\n")
}

func showAll() {
	fmt.Printf("=========== all ===========\n")
	for k := range all {
		fmt.Printf("%d ", k)
	}
	if len(all) == 0 {
		fmt.Printf("(null)")
	}
	fmt.Printf("\n")
}

func showInfo() {
	showAll()
	showNear()
	showCost()
}
