package graph

func Topology(g *Graph) ([]int64, bool) {
	queue := make([]int64, g.n, g.n)
	din := make([]int64, g.n, g.n)
	copy(din, g.din)

	hh, tt := int64(0), int64(-1)

	for i := int64(0); i < g.n; i++ {
		if din[i] == 0 {
			tt++
			queue[tt] = i
		}
	}

	for hh <= tt {
		t := queue[hh]
		hh++

		for i := g.h[t]; i != -1; i = g.ne[i] {
			j := g.e[i]
			din[j]--
			if din[j] == 0 {
				tt++
				queue[tt] = j
			}
		}
	}
	return queue, tt == g.n-1
}
