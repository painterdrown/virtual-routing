package router

// Test .
func Test() {
	port = 3001

	all[3001] = true
	all[3002] = true
	all[3003] = true
	all[3004] = true
	all[3005] = true

	near[3003] = true
	near[3004] = true
	near[3005] = true

	cost[3001] = make(map[int]int)
	cost[3002] = make(map[int]int)
	cost[3003] = make(map[int]int)
	cost[3004] = make(map[int]int)
	cost[3005] = make(map[int]int)

	cost[3001][3003] = 1
	cost[3001][3004] = 4
	cost[3001][3005] = 6

	cost[3002][3004] = 3
	cost[3002][3005] = 2

	cost[3003][3001] = 1
	cost[3003][3004] = 2

	cost[3004][3001] = 4
	cost[3004][3002] = 3
	cost[3004][3003] = 2

	cost[3005][3001] = 6
	cost[3005][3002] = 2

	updateRoutingTable()
}
