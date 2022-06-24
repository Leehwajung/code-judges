package main

import (
	"testing"
)

func Test_jobScheduling(t *testing.T) {
	type args struct {
		startTime []int
		endTime   []int
		profit    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				startTime: []int{1, 2, 3, 3},
				endTime:   []int{3, 4, 5, 6},
				profit:    []int{50, 10, 40, 70},
			},
			want: 120,
		},
		{
			name: "Example 2",
			args: args{
				startTime: []int{1, 2, 3, 4, 6},
				endTime:   []int{3, 5, 10, 6, 9},
				profit:    []int{20, 20, 100, 70, 60},
			},
			want: 150,
		},
		{
			name: "Example 3",
			args: args{
				startTime: []int{1, 1, 1},
				endTime:   []int{2, 3, 4},
				profit:    []int{5, 6, 4},
			},
			want: 6,
		},
		{
			name: "Case 9",
			args: args{
				startTime: []int{4, 2, 4, 8, 2},
				endTime:   []int{5, 5, 5, 10, 8},
				profit:    []int{1, 2, 8, 10, 4},
			},
			want: 18,
		},
		{
			name: "Case 29",
			args: args{
				startTime: []int{4, 3, 1, 2, 4, 8, 3, 3, 3, 9},
				endTime:   []int{5, 6, 3, 5, 9, 9, 8, 5, 7, 10},
				profit:    []int{9, 9, 5, 12, 10, 11, 10, 4, 14, 7},
			},
			want: 37,
		},
		{
			name: "Custom 1",
			args: args{
				startTime: []int{1, 2, 3, 4, 6},
				endTime:   []int{3, 5, 10, 6, 9},
				profit:    []int{20, 20, 140, 70, 60},
			},
			want: 160,
		},
		{
			name: "Custom 2",
			args: args{
				startTime: []int{1, 2, 3, 4, 6},
				endTime:   []int{3, 5, 10, 6, 9},
				profit:    []int{20, 20, 130, 70, 60},
			},
			want: 150,
		},
		{
			name: "Custom 3",
			args: args{
				startTime: []int{5, 2, 4, 7, 2},
				endTime:   []int{6, 5, 5, 9, 8},
				profit:    []int{1, 2, 8, 10, 4},
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jobScheduling(tt.args.startTime, tt.args.endTime, tt.args.profit); got != tt.want {
				t.Errorf("jobScheduling() = %v, want %v", got, tt.want)
			}
		})
	}
}
