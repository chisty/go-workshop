package main

import "fmt"

func main() {
	fmt.Println("Construct Binary Search Tree from Preorder Traversal")
	doTest([]int{4, 8, 12})
	doTest([]int{8, 5, 1, 7, 10, 12})
	doTest([]int{14, 15, 20, 22})
	doTest([]int{10, 8, 5, 3})
	doTest([]int{14, 10, 8, 5, 3, 15, 20, 22})
	doTest([]int{14, 10, 8, 5, 3, 4, 9, 15, 20, 16, 18, 22})
}

func doTest(input []int) {
	fmt.Println(input)
	res := bstFromPreorder(input)
	preorder(res)
	fmt.Println("\nDone")
}

func bstFromPreorder(preorder []int) *TreeNode {
	l := len(preorder)
	if l == 0 {
		return nil
	}

	root := constructTree(preorder, 0, l)
	return root
}

//Try to get pivot point to divide into left and right tree
func constructTree(preorder []int, start, l int) *TreeNode {
	if start == l {
		return nil
	}

	root := newNode(preorder[start])
	pivot := -1
	for i := start + 1; i < l; i++ {
		if preorder[i] > preorder[start] {
			pivot = i
			break
		}
	}

	if pivot == -1 {
		pivot = l
	}

	root.Left = constructTree(preorder, start+1, pivot)
	root.Right = constructTree(preorder, pivot, l)

	return root
}

//TreeNode data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func preorder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d,%v, %v\n", node.Val, node.Left, node.Right)
	preorder(node.Left)
	preorder(node.Right)
}
