package solution

import (
	"fmt"
	"sort"
)

func solution() {
	var t byte
	fmt.Scanf("%d", &t)

	var i byte
	for i = 1; i <= t; i++ {
		var n, m int
		fmt.Scanf("%d %d", &n, &m)

		a := make([]int, 0, n)
		for j := 0; j < n; j++ {
			var k int
			fmt.Scanf("%d", &k)
			a = append(a, k)
		}

		res := maxMetricValue(n, m, a)
		fmt.Printf("Case #%d: %f\n", i, res)
	}
}

func maxMetricValue(n, m int, a []int) float64 {
	var res float64 = 0

	sort.Ints(a)
	i := n - 1
	for ; m > 1; m-- {
		res += float64(a[i])
		i--
	}
	res += median(a, 0, i)

	return res
}

func median(sortedArray []int, start, end int) float64 {
	if (end-start)&1 == 1 {
		// even number of elements
		mid := ((end - start) >> 1)
		return float64(sortedArray[start+mid]+sortedArray[start+mid+1]) / 2
	}

	return float64(sortedArray[start+((end-start)>>1)])
}
