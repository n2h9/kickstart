package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var t byte
	fmt.Scanf("%d", &t)

	var i byte
	for i = 1; i <= t; i++ {
		var n int
		fmt.Scanf("%d\n", &n)

		a := make([]int, 0, n)
		for j := 0; j < n; j++ {
			var k int
			fmt.Scanf("%d", &k)
			a = append(a, k)
		}

		res := highestPossibleRating(a)
		fmt.Printf("Case #%d: %s\n", i, strings.Join(toString(res), " "))
	}
}

func highestPossibleRating(rates []int) []int {
	ratesSorted := make([]int, len(rates))
	copy(ratesSorted, rates)
	prates := make([]int, len(rates))
	sort.Ints(ratesSorted)
	count := make(map[int]int)
	for _, rate := range rates {
		count[rate]++
	}

	for i := 0; i < len(rates); i++ {
		maxRate := 2 * rates[i]
		mentorIndex := sort.SearchInts(ratesSorted, maxRate)
		if mentorIndex == len(ratesSorted) || ratesSorted[mentorIndex] != maxRate {
			mentorIndex--
		}
		if mentorIndex < 0 {
			prates[i] = -1
			continue
		}
		if count[rates[i]] > 1 {
			prates[i] = ratesSorted[mentorIndex]
			continue
		}
		for ; mentorIndex >= 0 && ratesSorted[mentorIndex] == rates[i]; mentorIndex-- {

		}
		if mentorIndex >= 0 {
			prates[i] = ratesSorted[mentorIndex]
			continue
		}
		prates[i] = -1
	}

	return prates
}

func toString(a []int) []string {
	s := make([]string, 0, len(a))
	for _, v := range a {
		s = append(s, fmt.Sprintf("%d", v))
	}
	return s
}
