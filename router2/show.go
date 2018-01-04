package router2

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
	for k1, v1 := range dist {
		fmt.Printf("[%d]: ", k1)
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)", k2, v2)
		}
		fmt.Printf("\n")
	}
	if len(dist) == 0 {
		fmt.Printf("(null)\n")
	}
}

func showPrev() {
	fmt.Printf("=========== prev ===========\n")
	for k1, v1 := range prev {
		fmt.Printf("[%d]: ", k1)
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)", k2, v2)
		}
		fmt.Printf("\n")
	}
	if len(prev) == 0 {
		fmt.Printf("(null)\n")
	}
}

func showNear() {
	fmt.Printf("=========== near ===========\n")
	for k1, v1 := range near {
		fmt.Printf("[%d]: ", k1)
		for v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Printf("\n")
	}
	if len(near) == 0 {
		fmt.Printf("(null)\n")
	}
}

func showAll() {
	fmt.Printf("=========== all ===========\n")
	for u := range all {
		fmt.Printf("%d ", u)
	}
	if len(all) == 0 {
		fmt.Printf("(null)")
	}
	fmt.Printf("\n")
}

func showInfo() {
	showDist()
	showPrev()
	showCost()
}
