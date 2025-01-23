package mergetripletstoformtargettriplet

func mergeTriplets(triplets [][]int, target []int) bool {
	validCount := make(map[int]bool)

	for _, triplet := range triplets {
		if triplet[0] > target[0] || triplet[1] > target[1] || triplet[2] > target[2] {
			continue
		}

		for i, v := range triplet {
			if v == target[i] {
				validCount[i] = true
			}
		}
	}

	return len(validCount) == 3
}
