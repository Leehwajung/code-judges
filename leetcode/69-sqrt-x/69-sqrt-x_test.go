package main

import (
	"testing"
)

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{x: 4}, want: 2},
		{name: "Example 2", args: args{x: 8}, want: 2},
		{name: "Custom 1", args: args{x: 1}, want: 1},
		{name: "Custom 2", args: args{x: 2}, want: 1},
		{name: "Custom 3", args: args{x: 3}, want: 1},
		{name: "Custom 4", args: args{x: 9}, want: 3},
		{name: "Custom 5", args: args{x: 13}, want: 3},
		{name: "Custom 6", args: args{x: 17}, want: 4},
		{name: "Custom 7", args: args{x: 23}, want: 4},
		{name: "Custom 8", args: args{x: 31}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
