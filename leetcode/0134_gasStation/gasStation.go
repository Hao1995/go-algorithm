package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	var ans int
	var total int
	var tank int
	for start := 0; start < len(gas); start++ {
		tank += gas[start] - cost[start]
		total += gas[start] - cost[start]
		if tank < 0 {
			ans = start + 1
			tank = 0
		}
	}

	if total < 0 {
		return -1
	}

	return ans
}
