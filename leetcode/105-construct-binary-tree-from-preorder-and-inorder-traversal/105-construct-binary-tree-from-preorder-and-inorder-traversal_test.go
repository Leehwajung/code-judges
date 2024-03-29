package main

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: intOrNilsToBinaryTree(3, 9, 20, nil, nil, 15, 7),
		},
		{
			name: "Example 2",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: intOrNilsToBinaryTree(-1),
		},
		{
			name: "Case 7",
			args: args{
				preorder: []int{1, 2},
				inorder:  []int{1, 2},
			},
			want: intOrNilsToBinaryTree(1, nil, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
