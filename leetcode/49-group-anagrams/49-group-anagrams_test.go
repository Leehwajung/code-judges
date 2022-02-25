package main

import (
	"reflect"
	"testing"

	"code-judges/internal/util"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Example 1",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "Example 2",
			args: args{strs: []string{""}},
			want: [][]string{{""}},
		},
		{
			name: "Example 3",
			args: args{strs: []string{"a"}},
			want: [][]string{{"a"}},
		},
		{
			name: "Case 39",
			args: args{strs: []string{"abbbbbbbbbbb", "aaaaaaaaaaab"}},
			want: [][]string{{"aaaaaaaaaaab"}, {"abbbbbbbbbbb"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !groupAnagramsEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func groupAnagramsEqual(x, y [][]string) bool {
	ySets := make([]map[string]bool, len(y))
	for i, ys := range y {
		ySets[i] = util.StringsToSet(ys)
	}
	for _, xs := range x {
		xSet := util.StringsToSet(xs)
		equal := false
		for _, ySet := range ySets {
			if reflect.DeepEqual(xSet, ySet) {
				equal = true
				break
			}
		}
		if !equal {
			return false
		}
	}
	return true
}
