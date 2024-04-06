package main

import (
	"slices"
	"testing"
)

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wants []string
	}{
		{
			name:  "Example 1",
			args:  args{s: "lee(t(c)o)de)"},
			wants: []string{"lee(t(c)o)de", "lee(t(co)de)", "lee(t(c)ode)"},
		},
		{
			name:  "Example 2",
			args:  args{s: "a)b(c)d"},
			wants: []string{"ab(c)d"},
		},
		{
			name:  "Example 3",
			args:  args{s: "))(("},
			wants: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.args.s); !slices.Contains(tt.wants, got) {
				t.Errorf("minRemoveToMakeValid() = %v, wants %v", got, tt.wants)
			}
		})
	}
}
