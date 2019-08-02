package tree

import (
	"fmt"
	"reflect"
)

type BSTNodeInterface interface {
	InitNode() *BSTNode
}
type BSTNode struct {
	Data  BSTNodeInterface
	Left  *BSTNode
	Right *BSTNode
}

type BST struct {
	Root        *BSTNode
	NodeCount   int
	BSTNodeType reflect.Type
	//should return true if v1 >= v2 else false
	Compare func(v1, v2 BSTNodeInterface) bool
}

func InitEmptyBST(nodeType reflect.Type, comparator func(v1, v2 BSTNodeInterface) bool) *BST {
	return &BST{nil, 0, nodeType, comparator}
}

func (bst *BST) IsEmpty() bool {
	return bst.Root == nil
}

func (bst *BST) InsertNode(node *BSTNode) {
	if bst.Root == nil {
		bst.Root = node
		return
	}
	bst.insertNodeUtil(bst.Root, node)
}
func (bst *BST) insertNodeUtil(root *BSTNode, node *BSTNode) {
	if root == nil {
		return
	}
	if bst.Compare(node.Data, root.Data) && root.Right == nil {
		root.Right = node
		return
	}
	if bst.Compare(root.Data, node.Data) && root.Left == nil {
		root.Left = node
		return
	}
	if bst.Compare(node.Data, root.Data) {
		bst.insertNodeUtil(root.Right, node)
	}
	if bst.Compare(root.Data, node.Data) {
		bst.insertNodeUtil(root.Left, node)
	}

}

func (bst *BST) PrintInOrderTraversal(root *BSTNode) {
	if root == nil {
		return
	}
	bst.PrintInOrderTraversal(root.Left)
	fmt.Println(root.Data )
	bst.PrintInOrderTraversal(root.Right)
}




