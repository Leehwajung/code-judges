package main

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, want: 4},
		{name: "Example 2", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3}, want: -1},
		{name: "Example 3", args: args{nums: []int{1}, target: 0}, want: -1},
		{name: "Custom 1", args: args{nums: []int{6, 7, 8, 0, 1, 2, 3, 4, 5}, target: 2}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
