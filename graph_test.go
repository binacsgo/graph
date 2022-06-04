package graph

import (
	"testing"
)

func Test_Graph(t *testing.T) {
	type args struct {
		n, m            int64
		es              [][]int64
		deleted         []int64
		leftAfterDelete int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal case",
			args: args{
				n: 4,
				m: 3,
				es: [][]int64{
					{0, 1},
					{0, 2},
					{0, 3},
				},
				deleted:         []int64{0, 1},
				leftAfterDelete: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.args.n, tt.args.m)
			for i := range tt.args.es {
				e := tt.args.es[i]
				g.AddEdge(e[0], e[1], 1)
			}
			g.DeleteEdge(tt.args.deleted...)
			if g.NumEdge() != tt.args.leftAfterDelete {
				t.Errorf("NumEdge() got = %v, want %v", g.NumEdge(), tt.args.leftAfterDelete)
			}
			for i := g.h[0]; i != -1; i = g.ne[i] {
				j := g.e[i]
				for _, id := range tt.args.deleted {
					if j == id {
						t.Errorf("Edge want to be deleted: %v", j)
					}
				}
			}
		})
	}
}
