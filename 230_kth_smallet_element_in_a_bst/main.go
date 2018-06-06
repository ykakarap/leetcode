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

	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 8,
			},
			Right: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 10,
				},
				Right: &TreeNode{
					Val: 12,
				},
			},
		},
	}

	for k := 1; k <= 9; k++ {
		fmt.Printf("%vth smallest element is : %v \n", k, kthSmallest(root, k))
	}

}

func kthSmallest(root *TreeNode, k int) int {
	leftCount := countTillLeaf(root.Left)
	if leftCount+1 == k {
		return root.Val
	}

	if k > leftCount+1 {
		// search on the right hand side
		return kthSmallest(root.Right, k-(leftCount+1))
	}
	return kthSmallest(root.Left, k)
}

func countTillLeaf(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		// at leaf
		return 1
	}

	count := 1
	if node.Left != nil {
		count += countTillLeaf(node.Left)
	}
	if node.Right != nil {
		count += countTillLeaf(node.Right)
	}

	return count
}
