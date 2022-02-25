package main

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "Example 2",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name: "Case 4",
			args: args{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{2, 3},
			},
			want: [][]int{{1, 5}},
		},
		{
			name: "Custom 1",
			args: args{
				intervals:   [][]int{{3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{2, 4},
			},
			want: [][]int{{2, 5}, {6, 7}, {8, 10}, {12, 16}},
		},
		{
			name: "Custom 2",
			args: args{
				intervals:   [][]int{{3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{2, 6},
			},
			want: [][]int{{2, 7}, {8, 10}, {12, 16}},
		},
		{
			name: "Custom 3",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 5},
			},
			want: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
		},
		{
			name: "Custom 4",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 10},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name: "Custom 5",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 19},
			},
			want: [][]int{{1, 2}, {3, 19}},
		},
		{
			name: "Custom 6",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 4}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{5, 6},
			},
			want: [][]int{{1, 2}, {3, 4}, {5, 7}, {8, 10}, {12, 16}},
		},
		{
			name: "Custom 7",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 4}, {7, 8}, {9, 10}, {12, 16}},
				newInterval: []int{5, 6},
			},
			want: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {12, 16}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
