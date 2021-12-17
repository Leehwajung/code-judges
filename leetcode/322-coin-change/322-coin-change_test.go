package main

import (
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{coins: []int{1, 2, 5}, amount: 11}, want: 3},
		{name: "Example 2", args: args{coins: []int{2}, amount: 3}, want: -1},
		{name: "Example 3", args: args{coins: []int{1}, amount: 0}, want: 0},
		{name: "Case 40", args: args{coins: []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, amount: 9864}, want: 24},
		{name: "Case 60", args: args{coins: []int{186, 419, 83, 408}, amount: 6249}, want: 20},
		{name: "Custom 1", args: args{coins: []int{127, 2, 50, 51}, amount: 128}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
