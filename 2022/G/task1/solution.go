package solution

type input struct {
	n      int
	m      int
	p      int
	scores [][]int
}

func minTotalAdditionalSteps(val *input) int {
	maxPerDayFunc := func(day int) int {
		max := 0
		for i := 0; i < val.m; i++ {
			if val.scores[i][day] > max {
				max = val.scores[i][day]
			}
		}
		return max
	}

	res := 0
	for i := 0; i < val.n; i++ {
		if v := maxPerDayFunc(i) - val.scores[val.p-1][i]; v > 0 {
			res += v
		}
	}

	return res
}
