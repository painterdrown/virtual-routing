package router3

import (
	"fmt"
)

func showDist() {
	fmt.Printf("=========== dist ===========\n")
	for k1, v1 := range dist {
		fmt.Printf("(%d,%d) ", k1, v1)
	}
	if len(dist) == 0 {
		fmt.Printf("(null)\n")
	} else {
		fmt.Printf("\n")
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

func showNext() {
	fmt.Printf("=========== next ===========\n")
	for k1, v1 := range next {
		fmt.Printf("(%d,%d) ", k1, v1)
	}
	if len(next) == 0 {
		fmt.Printf("(null)\n")
	} else {
		fmt.Printf("\n")
	}
}

func showInfo() {
	showAll()
	showNear()
	showDist()
	showNext()
}
