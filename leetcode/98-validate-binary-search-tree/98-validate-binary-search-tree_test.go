package main

import (
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root: intOrNilsToBinaryTree(2, 1, 3),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root: intOrNilsToBinaryTree(5, 1, 4, nil, nil, 3, 6),
			},
			want: false,
		},
		{
			name: "Case 69",
			args: args{
				root: intOrNilsToBinaryTree(2, 2, 2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
