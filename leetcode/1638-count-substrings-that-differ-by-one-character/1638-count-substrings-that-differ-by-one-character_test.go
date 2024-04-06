package main

import (
	"testing"
)

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "aba",
				t: "baba",
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				s: "ab",
				t: "bb",
			},
			want: 3,
		},
		{
			name: "Custom 1",
			args: args{
				s: "abbab",
				t: "bbbbb",
			},
			want: 33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
