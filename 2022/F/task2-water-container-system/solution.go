package solution

import (
	"sort"
)

func numOfComplitelyFilledContainers(n int, connections [][]int, queries []int) int {
	for _, c := range connections {
		sort.Ints(c)
	}
	sort.Slice(connections, func(i, j int) bool {
		li := connections[i][0]
		// if connections[i][1] < li {
		// 	li = connections[i][1]
		// }
		lj := connections[j][0]
		// if connections[j][1] < lj {
		// 	lj = connections[j][1]
		// }
		return li < lj
	})

	levelsCapacity := make([]int, n) // [level] -> number of containers per level === to water capacity
	levelsCapacity[0] = 1            // only one container per bottom level

	containerToLevel := make([]int, n+1) // container to level mapping (0 is not in use) (starts from 1)
	containerToLevel[1] = 0

	adjList := make([][]int, n+1) // [i] -> connected to list (started from 1)
	for _, c := range connections {
		adjList[c[0]] = append(adjList[c[0]], c[1])
		adjList[c[1]] = append(adjList[c[1]], c[0])
	}

	visited := make([]bool, n+1)
	q := newQueue()
	q.push(1)

	for !q.isEmpty() {
		c := q.pop()
		visited[c] = true

		l := containerToLevel[c] + 1

		children := adjList[c]
		for _, child := range children {
			if visited[child] {
				continue
			}
			containerToLevel[child] = l
			levelsCapacity[l]++
			q.push(child)
		}
	}

	// for _, c := range connections {
	// 	level := containerToLevel[c[0]] + 1
	// 	containerToLevel[c[1]] = level
	// 	levelsCapacity[level]++
	// }

	level := 0
	water := 0
	res := 0

	for range queries {
		water += 1
		if water < levelsCapacity[level] {
			continue
		}
		water = 0
		res += levelsCapacity[level]
		level++
	}

	return res
}

// ----
// queue

type queue []int

func newQueue() queue {
	return make([]int, 0, 32)
}

func (q *queue) isEmpty() bool {
	return len(*q) <= 0
}

func (q *queue) push(node int) {
	(*q) = append(*q, node)
}

// panic on empty queue
func (q *queue) pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

// end of queue
// ----
