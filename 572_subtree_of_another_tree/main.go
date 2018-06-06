package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	s := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
				// Left: &TreeNode{
				// 	Val: 0,
				// },
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	t := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	s2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 2,
						},
					},
				},
			},
		},
	}

	t2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
	}

	fmt.Printf("is subtree? %v \n", isSubtree(s, t))
	fmt.Printf("is subtree? %v \n", isSubtree(s2, t2))

}

func isSubtree(s *TreeNode, t *TreeNode) bool {

	if (s == nil || t == nil) && s != t {
		return false
	}

	if !isSameTree(s, t) {
		return isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}

	return true
}

func isSameTree(a *TreeNode, b *TreeNode) bool {
	if (a == nil || b == nil) && a != b {
		return false
	}

	if a == nil && b == nil {
		return true
	}

	if a.Val != b.Val {
		return false
	}

	if !isSameTree(a.Left, b.Left) {
		return false
	}

	if !isSameTree(a.Right, b.Right) {
		return false
	}

	return true

}
