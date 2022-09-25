package solution

import (
	"sort"
	"strings"
)

type fabric struct {
	color      string
	durability int
	id         int
}

func numOfFabricks(fabrics []*fabric) int {
	fabrics2 := make([]*fabric, len(fabrics))
	copy(fabrics2, fabrics)

	adaSortCompFunc := func(i, j int) bool {
		lexCompareRes := strings.Compare(fabrics[i].color, fabrics[j].color)
		if lexCompareRes == 0 {
			return fabrics[i].id < fabrics[j].id
		}

		return lexCompareRes < 0
	}

	charlesSortCompFunc := func(i, j int) bool {
		if fabrics2[i].durability == fabrics2[j].durability {
			return fabrics2[i].id < fabrics2[j].id
		}

		return fabrics2[i].durability < fabrics2[j].durability
	}

	sort.Slice(fabrics, adaSortCompFunc)
	sort.Slice(fabrics2, charlesSortCompFunc)

	res := 0

	for i := 0; i < len(fabrics); i++ {
		// pointer comarission is ok here
		if fabrics[i] == fabrics2[i] {
			res++
		}
	}

	return res
}
