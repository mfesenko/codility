package discintersections

import "sort"

const maxIntersectionsCount = 10000000

type (
	edge struct {
		x     int
		start bool
	}

	edges []edge
)

// Solution returns the number of pairs of intersecting discs. If the number of intersections exceeds 10,000,000, returns -1.
// Each i-th element of the array represents a disc with the center at (i, 0) and radius a[i].
// Disks i and j intersect if they have at least one point in common (i != j).
// Length of the array a is an integer in the range [0..100,000].
// Element of the array a are integers in the range [0..2,147,483,647].
func Solution(a []int) int {
	e := make(edges, 0, len(a)*2)
	for i, radius := range a {
		e = append(e, newEdge(i-radius, true), newEdge(i+radius, false))
	}
	sort.Sort(e)
	intersectionsCount := 0
	currentIntersections := 0
	for _, edge := range e {
		if edge.start {
			currentIntersections++
			continue
		}
		currentIntersections--
		intersectionsCount += currentIntersections
	}
	if intersectionsCount > maxIntersectionsCount {
		return -1
	}
	return intersectionsCount
}

func newEdge(x int, start bool) edge {
	return edge{
		x:     x,
		start: start,
	}
}

func (e edges) Len() int {
	return len(e)
}

func (e edges) Less(i, j int) bool {
	di := e[i]
	dj := e[j]
	if di.x == dj.x {
		return di.start
	}
	return di.x < dj.x
}

func (e edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
