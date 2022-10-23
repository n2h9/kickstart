package main

import (
	"fmt"
)

func getScore(arr []int) int {
	res := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			continue
		}
		s := 0
		for j := i; j < len(arr); j++ {
			s += arr[j]
			if s < 0 {
				break
			}
			res += s
		}
	}
	return res
}

func main() {
	var t byte
	fmt.Scanf("%d", &t)

	var i byte
	for i = 1; i <= t; i++ {
		var n int
		fmt.Scanf("%d\n", &n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &arr[j])
		}

		r := getScore(arr)
		fmt.Printf("Case #%d: %d\n", i, r)
	}
}
