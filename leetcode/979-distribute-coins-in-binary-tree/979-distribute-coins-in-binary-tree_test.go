package main

import (
	"testing"
)

func Test_distributeCoins(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root: intOrNilsToBinaryTree(3, 0, 0),
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				root: intOrNilsToBinaryTree(0, 3, 0),
			},
			want: 3,
		},
		{
			name: "Custom 1",
			args: args{
				root: intOrNilsToBinaryTree(0, 0, 3),
			},
			want: 3,
		},
		{
			name: "Custom 2",
			args: args{
				root: intOrNilsToBinaryTree(1, 0, nil, 3, 0),
			},
			want: 3,
		},
		{
			name: "Custom 3",
			args: args{
				root: intOrNilsToBinaryTree(1, 0, nil, 0, 3),
			},
			want: 3,
		},
		{
			name: "Custom 4",
			args: args{
				root: intOrNilsToBinaryTree(1, 1, 1),
			},
			want: 0,
		},
		{
			name: "Custom 5",
			args: args{
				root: intOrNilsToBinaryTree(1, 2, 3, 0, 0, 0),
			},
			want: 5,
		},
		{
			name: "Custom 6",
			args: args{
				root: intOrNilsToBinaryTree(0, 0, 2, 4, nil, 0, 0),
			},
			want: 8,
		},
		{
			name: "Custom 7",
			args: args{
				root: intOrNilsToBinaryTree(0, 0, 0, 4, nil, 1, 1),
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCoins(tt.args.root); got != tt.want {
				t.Errorf("distributeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
