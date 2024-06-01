package main

import (
	"maps"
	"testing"

	"code-judges/internal/typeconv"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			want: []string{"cats and dog", "cat sand dog"},
		},
		{
			name: "Example 2",
			args: args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			want: []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			name: "Example 3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); !maps.Equal(typeconv.StringsToSet(got), typeconv.StringsToSet(tt.want)) {
				t.Errorf("wordBreak() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
