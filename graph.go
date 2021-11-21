package graph

type Graph struct {
	h, e, w, ne, din []int64
	n, m, tot        int64
}

func NewGraph(n, m int64) *Graph {
	g := &Graph{
		h:   make([]int64, n, n),
		e:   make([]int64, m, m),
		w:   make([]int64, m, m),
		ne:  make([]int64, m, m),
		din: make([]int64, n, n),
		n:   n,
		m:   m,
		tot: 0,
	}
	for i := range g.h {
		g.h[i] = -1
	}
	return g
}

func (g *Graph) AddEdge(a, b, c int64) {
	g.e[g.tot], g.w[g.tot], g.ne[g.tot], g.h[a] = b, c, g.h[a], g.tot
	g.tot++
	g.din[b]++
}
