package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{head: intsToLinkedList(1, 2, 3, 4, 5)},
			want: intsToLinkedList(5, 4, 3, 2, 1),
		},
		{
			name: "Example 2",
			args: args{head: intsToLinkedList(1, 2)},
			want: intsToLinkedList(2, 1),
		},
		{
			name: "Example 3",
			args: args{head: intsToLinkedList()},
			want: intsToLinkedList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
