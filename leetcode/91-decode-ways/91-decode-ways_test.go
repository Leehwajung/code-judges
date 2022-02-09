package main

import (
	"testing"
)

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example 1", args: args{s: "12"}, want: 2},
		{name: "Example 2", args: args{s: "226"}, want: 3},
		{name: "Example 3", args: args{s: "06"}, want: 0},
		{name: "Case 249", args: args{s: "10"}, want: 1},
		{name: "Case 263", args: args{s: "2101"}, want: 1},
		{name: "Custom 1", args: args{s: "22622"}, want: 6},
		{name: "Custom 2", args: args{s: "28622"}, want: 2},
		{name: "Custom 3", args: args{s: "18622"}, want: 4},
		{name: "Custom 4", args: args{s: "0"}, want: 0},
		{name: "Custom 5", args: args{s: "0000000000"}, want: 0},
		{name: "Custom 6", args: args{s: "0000010000"}, want: 0},
		{name: "Custom 7", args: args{s: "1010101010"}, want: 1},
		{name: "Custom 8", args: args{s: "10101010100"}, want: 0},
		{name: "Custom 9", args: args{s: "1234567891011121314151617181920212223242526"}, want: 259584},
		{name: "Custom 10", args: args{s: "12345678910111213141501617181920212223242526"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
