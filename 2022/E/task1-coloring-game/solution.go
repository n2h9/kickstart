package main

import "fmt"

func main() {
	var t byte
	fmt.Scanf("%d", &t)

	var i byte
	for i = 1; i <= t; i++ {
		var n int
		fmt.Scanf("%d\n", &n)

		res := coloringGame(n)
		fmt.Printf("Case #%d: %d\n", i, res)
	}
}

func coloringGame(n int) int {
	return (n-1)/5 + 1
}
