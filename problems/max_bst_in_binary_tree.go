package problems

import (
	"go-google/utils"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type ReturnValue struct {
	maxInSubTree int
	minInSubTree int
	nodesInSubTree int
	isSubTreeABst bool
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
var maxBstSize  = -1
func maxBst(root  *TreeNode) ReturnValue {
	if root == nil {
		maxBstSize = utils.Max(maxBstSize, 0)
		return ReturnValue{100000, -100000, 0, true}
	}
	if isLeaf(root) {
		maxBstSize = utils.Max(maxBstSize, 1)
		return ReturnValue{root.Data, root.Data, 1, true}
	}
	leftRetVal := maxBst(root.Left)
	rightRetVal := maxBst(root.Right)
	//merging the two bsts
	if leftRetVal.isSubTreeABst && rightRetVal.isSubTreeABst && root.Data > leftRetVal.maxInSubTree && root.Data < rightRetVal.minInSubTree {
		size := 1 + leftRetVal.nodesInSubTree + rightRetVal.nodesInSubTree
		maxBstSize = utils.Max(maxBstSize, size)
		return ReturnValue{rightRetVal.maxInSubTree, leftRetVal.minInSubTree, size, true }
	}
	if !leftRetVal.isSubTreeABst && !rightRetVal.isSubTreeABst {
		maxBstSize = utils.Max(maxBstSize, 1)
		return ReturnValue{root.Data, root.Data, 1, true}
	}
	if leftRetVal.isSubTreeABst && root.Data >= leftRetVal.maxInSubTree{
		maxBstSize = utils.Max(maxBstSize, 1+leftRetVal.nodesInSubTree)
		return ReturnValue{utils.Max(leftRetVal.maxInSubTree, root.Data), leftRetVal.minInSubTree, 1+leftRetVal.nodesInSubTree, true}
	}

	if rightRetVal.isSubTreeABst && root.Data <= rightRetVal.minInSubTree {
		maxBstSize = utils.Max(maxBstSize, 1+rightRetVal.nodesInSubTree)
		return ReturnValue{rightRetVal.maxInSubTree, utils.Min(root.Data, rightRetVal.minInSubTree), 1+rightRetVal.nodesInSubTree, true}
	}
	return ReturnValue{root.Data, root.Data, 1, true}



}

func FindMaxSizeBSTInBinaryTree(root *TreeNode) int{
	maxBst(root)
	return maxBstSize
}


