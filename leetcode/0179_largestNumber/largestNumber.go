package largestnumber

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	// convert to string
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}

	// compare two string
	sort.Slice(strs, func(i, j int) bool {
		var stri, strj string = strs[i], strs[j]
		return (stri + strj) > (strj + stri)
	})

	// join str
	var b strings.Builder
	for _, str := range strs {
		b.WriteString(str)
	}

	ans := b.String()

	// handle edge case
	if val, _ := strconv.Atoi(ans); val == 0 {
		return "0"
	}

	return ans
}
