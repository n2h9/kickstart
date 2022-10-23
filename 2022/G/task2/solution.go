package main

import (
	"fmt"
	"sort"
)

type input struct {
	rs      int
	rh      int
	n       int
	ncoords [][]int
	m       int
	mcoords [][]int
}

func getScore(inp *input) (int, int) {
	sqr := func(x int) int {
		return x * x
	}

	// radius effectvie square
	// radius which allows to store to be treated as at home
	// to not calculate the root use square of the radius
	re := sqr(inp.rh + inp.rs)

	type tteam byte
	const TEAM_READ tteam = 0
	const TEAM_YELLOW tteam = 1

	type stone struct {
		t tteam // team
		d int   // distance from the center
	}

	// do not take a root, leave as a square
	// it fits the int
	dist := func(x, y int) int {
		return sqr(x) + sqr(y)
	}

	isAtHome := func(d int) bool {
		return (d - re) <= 0
	}

	stones := make([]*stone, 0, inp.n+inp.m)

	// red team stones
	for _, s := range inp.ncoords {
		d := dist(s[0], s[1])
		if isAtHome(d) {
			stones = append(stones, &stone{
				t: TEAM_READ,
				d: d,
			})
		}
	}

	// yellow team stones
	for _, s := range inp.mcoords {
		d := dist(s[0], s[1])
		if isAtHome(d) {
			stones = append(stones, &stone{
				t: TEAM_YELLOW,
				d: d,
			})
		}
	}

	if len(stones) <= 0 {
		return 0, 0
	}

	sort.Slice(stones, func(i, j int) bool {
		return stones[i].d <= stones[j].d
	})

	redWin := stones[0].t == TEAM_READ
	res := 1

	for i := 1; i < len(stones) && stones[i].t == stones[i-1].t; i++ {
		res++
	}

	if redWin {
		return res, 0
	}
	return 0, res
}

func main() {
	var t byte
	fmt.Scanf("%d", &t)

	var i byte
	for i = 1; i <= t; i++ {
		var rs, rh, n, m int
		fmt.Scanf("%d %d\n", &rs, &rh)
		fmt.Scanf("%d\n", &n)
		nscores := make([][]int, n)
		for j := 0; j < n; j++ {
			nscores[j] = make([]int, 2)
			fmt.Scanf("%d %d\n", &nscores[j][0], &nscores[j][1])
		}
		fmt.Scanf("%d\n", &m)
		mscores := make([][]int, m)
		for j := 0; j < m; j++ {
			mscores[j] = make([]int, 2)
			fmt.Scanf("%d %d\n", &mscores[j][0], &mscores[j][1])
		}

		r, y := getScore(
			&input{
				rs:      rs,
				rh:      rh,
				n:       n,
				ncoords: nscores,
				m:       m,
				mcoords: mscores,
			},
		)
		fmt.Printf("Case #%d: %d %d\n", i, r, y)
	}
}
