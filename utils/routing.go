package utils

import (
	"strconv"
	"strings"

	"github.com/painterdrown/virtual-routing/global"
)

// UpdateCost .
func UpdateCost(source string, costs []string) bool {
	var needToUpdateRoutingTable = false
	for _, v := range costs {
		parts := strings.Split(v, " ")
		dest := parts[0]
		cost, _ := strconv.Atoi(parts[1])
		if global.Cost[source] != nil && global.Cost[source][dest] != cost {
			needToUpdateRoutingTable = true
			global.Cost[source][dest] = cost
		}
	}
	return needToUpdateRoutingTable
}

// UpdateRoutingTable .
func UpdateRoutingTable() {

}
