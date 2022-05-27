package main

import (
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{intervals: [][]int{{1, 2}, {1, 2}, {1, 2}}},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{intervals: [][]int{{1, 2}, {2, 3}}},
			want: 0,
		},
		{
			name: "Case 19",
			args: args{intervals: [][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}},
			want: 7,
		},
		{
			name: "Custom 1",
			args: args{intervals: [][]int{{27, 74}, {35, 89}, {37, 102}, {38, 51}, {48, 131}, {60, 74}, {69, 149}, {130, 147}, {158, 195}, {166, 198}, {182, 197}, {195, 199}}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
