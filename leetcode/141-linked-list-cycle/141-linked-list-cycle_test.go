package main

import (
	"testing"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{head: linkedListWithCycle(1, 3, 2, 0, -4)},
			want: true,
		},
		{
			name: "Example 2",
			args: args{head: linkedListWithCycle(0, 1, 2)},
			want: true,
		},
		{
			name: "Example 3",
			args: args{head: linkedListWithCycle(-1, 0)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
