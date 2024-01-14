package lettercombinationsofaphonenumber

var (
	digitStrMap = make(map[string][]string)
)

func init() {
	digitStrMap["2"] = []string{"a", "b", "c"}
	digitStrMap["3"] = []string{"d", "e", "f"}
	digitStrMap["4"] = []string{"g", "h", "i"}
	digitStrMap["5"] = []string{"j", "k", "l"}
	digitStrMap["6"] = []string{"m", "n", "o"}
	digitStrMap["7"] = []string{"p", "q", "r", "s"}
	digitStrMap["8"] = []string{"t", "u", "v"}
	digitStrMap["9"] = []string{"w", "x", "y", "z"}
}

func letterCombinations(digits string) []string {
	var cands []string
	for i := 0; i < len(digits); i++ {
		cands = append(cands, string(digits[i]))
	}

	var ans []string
	backtracking(&ans, cands, "")
	return ans
}

func backtracking(ans *[]string, cands []string, tmp string) {
	if len(cands) == 0 {
		if tmp != "" {
			*ans = append(*ans, tmp)
		}
		return
	}

	cand := cands[0]
	for _, candStr := range digitStrMap[cand] {
		tmp += candStr
		backtracking(ans, cands[1:], tmp)
		tmp = tmp[:len(tmp)-1]
	}
}
