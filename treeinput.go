package goleet

type TreeNode struct {
	Val int
	Right *TreeNode
	Left *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func tree(arr []*int, root *TreeNode, i int ) *TreeNode {
	if i < len(arr) && arr[i] != nil {
		temp := NewTreeNode(*arr[i])
		root = temp

		root.Left = tree(arr,root.Left, 2*i+1)
		root.Right = tree(arr,root.Right, 2*i+2)
	}
	return root
	 
}

func Tree(arr []*int) *TreeNode {
	var root *TreeNode
	return tree(arr, root, 0)
}