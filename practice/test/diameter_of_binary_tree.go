package main

import (
	"fmt"
)

func main() {
	fmt.Println("Diameter of Binary Tree")
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 5}},
			Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 6, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 8}}},
		},
		Right: &TreeNode{
			Val:   10,
			Left:  &TreeNode{Val: 11, Left: &TreeNode{Val: 12}},
			Right: &TreeNode{Val: 13, Left: &TreeNode{Val: 14, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 16}}},
		},
	}

	dia := diameterOfBinaryTree(root)
	fmt.Println("Result= ", dia)
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, dia := traverseTree(root)
	//fmt.Printf("val= %d, dia= %d\n", val, dia)
	return dia
}

func traverseTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	lh, ld := traverseTree(root.Left)
	rh, rd := traverseTree(root.Right)

	//fmt.Printf("Current= %d, lh=%d, ld=%d. rh=%d, rd=%d\n", root.Val, lh, ld, rh, rd)
	return getMax(lh, rh) + 1, getMax(lh+rh, getMax(ld, rd))
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
