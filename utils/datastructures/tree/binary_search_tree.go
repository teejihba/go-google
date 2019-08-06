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

var inOrderSuccessor *BSTNode


//Returns nil if there is no successor
func (bst *BST) InOrderSuccessor(root, node *BSTNode) *BSTNode {
	inOrderSuccessor = nil
	if root == nil {
		return nil
	}
	bst.inOrderSuccessorUtil(root, node)
	return inOrderSuccessor
}
func (bst *BST) inOrderSuccessorUtil(root, node *BSTNode) {
	if root == nil {
		return
	}
	if root == node {
		if root.Right != nil {
			bst.findMin(root)
		}
		return
	} else if bst.Compare(root.Data, node.Data) {
		inOrderSuccessor = root
		bst.inOrderSuccessorUtil(root.Left, node)
	} else {
		bst.inOrderSuccessorUtil(root.Right, node)
	}
}

func (bst *BST) findMin(root *BSTNode) {
	temp := root.Right
	for temp != nil {
		inOrderSuccessor = temp
		temp = temp.Left
	}
}
var kSuccessors []*BSTNode
func (bst *BST) KInOrderSuccessors(root, node *BSTNode, k int) []*BSTNode{
	kSuccessors = []*BSTNode{}
	bst.kSuccessorsUtil(root,node,k)
	return kSuccessors
}

func(bst *BST) kSuccessorsUtil(root, node *BSTNode, k int){
	if root == nil {
		return
	}
	if root == node {
		if root.Right != nil {
			bst.inOrderAddNodesInListUtil(root.Right, k)
		}
	} else if bst.Compare(root.Data, node.Data) {
		bst.kSuccessorsUtil(root.Left, node, k)
		kSuccessors = append(kSuccessors, root)
		bst.inOrderAddNodesInListUtil(root.Right, k)
	} else {
		bst.kSuccessorsUtil(root.Right, node, k)
	}
}

func (bst *BST) inOrderAddNodesInListUtil(node * BSTNode, k int){
	if len(kSuccessors) >= k || node == nil {
		return
	}
	bst.inOrderAddNodesInListUtil(node.Left, k)
	if len(kSuccessors) >= k {
		return
	}
	kSuccessors = append(kSuccessors, node)
	bst.inOrderAddNodesInListUtil(node.Right, k)
}




