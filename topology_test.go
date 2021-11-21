package graph

import (
	"reflect"
	"testing"
)

func Test_Topology(t *testing.T) {
	type args struct {
		n, m int64
		es   [][]int64
	}
	tests := []struct {
		name  string
		args  args
		want  []int64
		want1 bool
	}{
		{
			name: "normal case 1",
			args: args{
				n: 4,
				m: 5,
				es: [][]int64{
					{0, 1},
					{1, 2},
					{0, 3},
					{1, 3},
					{2, 3},
				},
			},
			want:  []int64{0, 1, 2, 3},
			want1: true,
		},
		{
			name: "normal case 2",
			args: args{
				n: 4,
				m: 5,
				es: [][]int64{
					{0, 1},
					{1, 2},
					{0, 3},
					{1, 3},
					{3, 2}, // inverse edge
				},
			},
			want:  []int64{0, 1, 3, 2},
			want1: true,
		},
		{
			name: "normal case 2",
			args: args{
				n: 4,
				m: 10,
				es: [][]int64{
					{0, 1},
					{1, 2},
					{0, 3},
					{1, 3},
					{2, 3},
					// multi edges
					{0, 1},
					{1, 2},
					{0, 3},
					{1, 3},
					{2, 3},
				},
			},
			want:  []int64{0, 1, 2, 3},
			want1: true,
		},
		{
			name: "circle",
			args: args{
				n: 4,
				m: 5,
				es: [][]int64{
					{0, 1},
					{1, 2},
					{0, 3},
					{1, 3},
					{3, 1}, // 1->3 3->1 circle
				},
			},
			want:  []int64{0, 0, 0, 0},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.args.n, tt.args.m)
			for i := range tt.args.es {
				e := tt.args.es[i]
				g.AddEdge(e[0], e[1], 1)
			}
			got, got1 := Topology(g)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Topology() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Topology() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
