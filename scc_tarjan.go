package graph

type SCC struct {
	dfn, low           []int64
	id, size           []int64
	timestamp, scc_cnt int64
}

func newSCC(g *Graph) *SCC {
	return &SCC{
		dfn:       make([]int64, g.n),
		low:       make([]int64, g.n),
		id:        make([]int64, g.n),
		size:      make([]int64, g.n),
		timestamp: 0,
		scc_cnt:   0,
	}
}

func SCCTarjan(g *Graph) *SCC {
	scc := newSCC(g)

	stk := make([]int64, g.n+1)
	in_stk := make([]bool, g.n+1)
	top := 0

	var tarjan func(u int64)
	tarjan = func(u int64) {
		{
			scc.timestamp++
			scc.dfn[u], scc.low[u] = scc.timestamp, scc.timestamp
		}
		{
			top++
			stk[top], in_stk[u] = u, true
		}
		for i := g.h[u]; i != -1; i = g.ne[i] {
			j := g.e[i]
			if scc.dfn[j] == 0 {
				tarjan(j)
				if scc.low[j] < scc.low[u] {
					scc.low[u] = scc.low[j]
				}
			} else if in_stk[j] {
				if scc.dfn[j] < scc.low[u] {
					scc.low[u] = scc.dfn[j]
				}
			}
		}
		if scc.dfn[u] == scc.low[u] {
			for {
				y := stk[top]
				top--
				in_stk[y] = false
				scc.id[y] = scc.scc_cnt
				scc.size[scc.scc_cnt]++
				if y == u {
					break
				}
			}
			scc.scc_cnt++
		}
	}

	for i := int64(0); i < g.n; i++ {
		if scc.dfn[i] == 0 {
			tarjan(i)
		}
	}
	return scc
}
