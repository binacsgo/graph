package graph

type Graph struct {
	h, e, w, ne, din   []int64
	fr, pre            []int64
	n, m, tot, deleted int64
}

func NewGraph(n, m int64) *Graph {
	g := &Graph{
		h:       make([]int64, n, n),
		e:       make([]int64, m, m),
		w:       make([]int64, m, m),
		ne:      make([]int64, m, m),
		din:     make([]int64, n, n),
		fr:      make([]int64, m, m),
		pre:     make([]int64, m, m),
		n:       n,
		m:       m,
		tot:     0,
		deleted: 0,
	}
	for i := range g.h {
		g.h[i] = -1
	}
	for i := range g.e {
		g.pre[i] = -1
	}
	return g
}

func (g *Graph) AddEdge(a, b, c int64) {
	if g.h[a] >= 0 {
		g.pre[g.h[a]] = g.tot
	}
	g.fr[g.tot] = a
	g.e[g.tot], g.w[g.tot], g.ne[g.tot], g.h[a] = b, c, g.h[a], g.tot
	g.din[b]++
	g.tot++
}

func (g *Graph) DeleteEdge(ids ...int64) {
	for _, id := range ids {
		if id < 0 || id >= g.tot {
			continue
		}
		if g.pre[id] >= 0 {
			g.ne[g.pre[id]] = g.ne[id]
		} else {
			// This means the current id is head.
			g.h[g.fr[id]] = g.ne[id]
		}
		if g.ne[id] >= 0 {
			g.pre[g.ne[id]] = g.pre[id]
		} else {
			// Do nothing.
		}
		g.deleted++
	}
}

func (g *Graph) NumEdge() int64 {
	return g.tot - g.deleted
}
