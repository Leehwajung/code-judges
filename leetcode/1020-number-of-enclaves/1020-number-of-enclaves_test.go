package main

import (
	"testing"
)

func Test_numEnclaves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				grid: [][]int{
					{0, 0, 0, 0},
					{1, 0, 1, 0},
					{0, 1, 1, 0},
					{0, 0, 0, 0},
				},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				grid: [][]int{
					{0, 1, 1, 0},
					{0, 0, 1, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 0},
				},
			},
			want: 0,
		},
		{
			name: "Case 17",
			args: args{
				grid: [][]int{
					{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
					{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
					{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
					{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
					{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
					{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
					{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
					{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
					{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 1, 0, 0, 0, 1},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEnclaves(tt.args.grid); got != tt.want {
				t.Errorf("numEnclaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
