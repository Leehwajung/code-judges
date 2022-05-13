package main

import (
	"reflect"
	"testing"
)

func Test_pacificAtlantic(t *testing.T) {
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				heights: [][]int{
					{1, 2, 2, 3, 5},
					{3, 2, 3, 4, 4},
					{2, 4, 5, 3, 1},
					{6, 7, 1, 4, 5},
					{5, 1, 1, 2, 4},
				},
			},
			want: [][]int{
				{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
			},
		},
		{
			name: "Example 2",
			args: args{
				heights: [][]int{
					{2, 1},
					{1, 2},
				},
			},
			want: [][]int{
				{0, 0}, {0, 1}, {1, 0}, {1, 1},
			},
		},
		{
			name: "Case 51",
			args: args{
				heights: [][]int{
					{1, 2, 3},
					{8, 9, 4},
					{7, 6, 5},
				},
			},
			want: [][]int{
				{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2},
			},
		},
		{
			name: "Custom 1",
			args: args{
				heights: [][]int{
					{8, 8, 8, 8, 9},
					{8, 1, 1, 1, 4},
					{8, 1, 5, 4, 1},
					{8, 1, 1, 4, 5},
					{9, 2, 2, 2, 4},
				},
			},
			want: [][]int{
				{0, 4}, {4, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pacificAtlantic(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pacificAtlantic() = %v, want %v", got, tt.want)
			}
		})
	}
}
