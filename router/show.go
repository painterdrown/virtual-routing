package router

import (
	"fmt"
)

// ShowCost .
func ShowCost() {
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

// ShowDist .
func ShowDist() {
	fmt.Printf("=========== dist ===========\n")
	for k, v := range dist {
		fmt.Printf("(%d,%d)", k, v)
	}
	if len(dist) == 0 {
		fmt.Printf("(null)")
	}
	fmt.Printf("\n")
}

// ShowPrev .
func ShowPrev() {
	fmt.Printf("=========== prev ===========\n")
	for k, v := range prev {
		fmt.Printf("(%d,%d)", k, v)
	}
	if len(prev) == 0 {
		fmt.Printf("(null)")
	}
	fmt.Printf("\n")
}

// ShowNear .
func ShowNear() {
	fmt.Printf("=========== near ===========\n")
	for k := range near {
		fmt.Printf("%d ", k)
	}
	if len(near) == 0 {
		fmt.Printf("(null)")
	}
	fmt.Printf("\n")
}

// ShowAll .
func ShowAll() {
	fmt.Printf("=========== all ===========\n")
	for k := range all {
		fmt.Printf("%d ", k)
	}
	if len(all) == 0 {
		fmt.Printf("(null)")
	}
	fmt.Printf("\n")
}

// ShowInfo .
func ShowInfo() {
	ShowNear()
	ShowDist()
	ShowPrev()
	ShowCost()
}
