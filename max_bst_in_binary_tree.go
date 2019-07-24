package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

type ReturnValue struct {
	maxInSubTree int
	minInSubTree int
	nodesInSubTree int
	isSubTreeABst bool
}

func isLeaf(node *Node) bool {
	return node.left == nil && node.right == nil
}
var maxBstSize  = -1
func maxBst(root  *Node) ReturnValue {
	if root == nil {
		maxBstSize = max(maxBstSize, 0)
		return ReturnValue{100000, -100000, 0, true}
	}
	if isLeaf(root) {
		maxBstSize = max(maxBstSize, 1)
		return ReturnValue{root.data, root.data, 1, true}
	}
	leftRetVal := maxBst(root.left)
	rightRetVal := maxBst(root.right)
	//merging the two bsts
	if leftRetVal.isSubTreeABst && rightRetVal.isSubTreeABst && root.data > leftRetVal.maxInSubTree && root.data < rightRetVal.minInSubTree {
		size := 1 + leftRetVal.nodesInSubTree + rightRetVal.nodesInSubTree
		maxBstSize = max(maxBstSize, size)
		return ReturnValue{rightRetVal.maxInSubTree, leftRetVal.minInSubTree, size, true }
	}
	if !leftRetVal.isSubTreeABst && !rightRetVal.isSubTreeABst {
		maxBstSize = max(maxBstSize, 1)
		return ReturnValue{root.data, root.data, 1, true}
	}
	if leftRetVal.isSubTreeABst && root.data >= leftRetVal.maxInSubTree{
		maxBstSize = max(maxBstSize, 1+leftRetVal.nodesInSubTree)
		return ReturnValue{max(leftRetVal.maxInSubTree, root.data), leftRetVal.minInSubTree, 1+leftRetVal.nodesInSubTree, true}
	}

	if rightRetVal.isSubTreeABst && root.data <= rightRetVal.minInSubTree {
		maxBstSize = max(maxBstSize, 1+rightRetVal.nodesInSubTree)
		return ReturnValue{rightRetVal.maxInSubTree, min(root.data, rightRetVal.minInSubTree), 1+rightRetVal.nodesInSubTree, true}
	}
	return ReturnValue{root.data, root.data, 1, true}



}

func max(i1 int, i2 int) int {
	if i1>=i2 {
		return i1
	}
	return i2
}
func min(i1 int, i2 int) int {
	if i1>=i2 {
		return i2
	}
	return i1
}

func main(){
	root := &Node{10, nil, nil}
	l := &Node{15, nil, nil}
	root.left = l
	r := &Node{15, nil, nil}
	root.right = r
	rr := &Node{20, nil,  nil}
	r.right = rr
	maxBst(root)
	fmt.Println("max bst size = ", maxBstSize)
}

