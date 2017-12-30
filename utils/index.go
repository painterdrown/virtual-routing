package utils

import (
	"fmt"
	"strconv"

	"github.com/painterdrown/virtual-routing/global"
)

// CheckErr .
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Prompt .
func Prompt(msg string) {
	fmt.Printf("\n%s\n[%d] > ", msg, global.Port)
}

// ShowCost .
func ShowCost() {
	for k1, v1 := range global.Cost {
		print("[" + strconv.Itoa(k1) + "]: ")
		for k2, v2 := range v1 {
			print("(" + strconv.Itoa(k2) + "," + strconv.Itoa(v2) + ")")
		}
		print("\n")
	}
}
