package tests

import (
	"go-google/problems"
	"testing"
)

func TestFindMaxSizeBSTInBinaryTree(t *testing.T) {
	root := &problems.TreeNode{10, nil, nil}
	l := &problems.TreeNode{15, nil, nil}
	root.Left = l
	r := &problems.TreeNode{15, nil, nil}
	root.Right = r
	rr := &problems.TreeNode{20, nil, nil}
	r.Right = rr
	expected := 3
	ans := problems.FindMaxSizeBSTInBinaryTree(root)
	if ans != expected {
		t.Errorf("test case failed, expected = %v, actual = %v", expected, ans)
	}
}
