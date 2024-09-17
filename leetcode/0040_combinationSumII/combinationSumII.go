package combinationsumii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var combs [][]int
	bt(candidates, target, 0, 0, []int{}, &combs)
	return combs
}

func bt(cands []int, target int, idx int, sum int, comb []int, combs *[][]int) {
	if sum == target {
		ncomb := make([]int, len(comb))
		copy(ncomb, comb)
		*combs = append(*combs, ncomb)
		return
	}

	if sum > target || idx == len(cands) {
		return
	}

	// include
	bt(cands, target, idx+1, sum+cands[idx], append(comb, cands[idx]), combs)

	// not include
	for idx+1 < len(cands) && cands[idx] == cands[idx+1] {
		idx += 1
	}
	bt(cands, target, idx+1, sum, comb, combs)

	return
}
