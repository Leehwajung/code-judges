package main

import (
	"testing"
)

func Test_numSubarraysWithSum(t *testing.T) {
	type args struct {
		nums []int
		goal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{nums: []int{1, 0, 1, 0, 1}, goal: 2}, want: 4},
		{name: "Example 2", args: args{nums: []int{0, 0, 0, 0, 0}, goal: 0}, want: 15},
		{name: "Custom 1", args: args{nums: []int{1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}, goal: 0}, want: 22},
		{name: "Custom 2", args: args{nums: []int{1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}, goal: 1}, want: 41},
		{name: "Custom 3", args: args{nums: []int{1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}, goal: 2}, want: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarraysWithSum(tt.args.nums, tt.args.goal); got != tt.want {
				t.Errorf("numSubarraysWithSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
