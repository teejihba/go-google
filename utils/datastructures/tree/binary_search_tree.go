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
	Difference func(v1, v2 BSTNodeInterface) int
}

func InitEmptyBST(nodeType reflect.Type, comparator func(v1, v2 BSTNodeInterface) bool ,
	diff func(v1,v2 BSTNodeInterface)int) *BST {
	return &BST{BSTNodeType: nodeType, Compare: comparator, Difference: diff}
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
var kPredecessors []*BSTNode
func (bst *BST) KInOrderSuccessors(root, node *BSTNode, k int) []*BSTNode{
	kSuccessors = []*BSTNode{}
	bst.kSuccessorsUtil(root,node,k)
	return kSuccessors
}

func (bst *BST) KInOrderPredecessors(root, node *BSTNode, k int) []*BSTNode{
	kPredecessors = []*BSTNode{}
	bst.kPredecessorsUtil(root,node,k)
	return kPredecessors
}

func (bst *BST) KClosestNodeOfANode(root , node *BSTNode, k int) []*BSTNode {
	kSuccessors := bst.KInOrderSuccessors(root, node, k)
	kPredecessors := bst.KInOrderPredecessors(root, node, k)
	var kClosest []*BSTNode
	i,j := 0,0
	sLen, pLen := len(kSuccessors), len(kPredecessors)
	var sVal , pVal *BSTNode
	for {
		if len(kClosest) == k {
			return kClosest
		}

		sVal , pVal = nil, nil
		if i < sLen {
			sVal = kSuccessors[i]
		}
		if j < pLen {
			pVal = kPredecessors[j]
		}

		if sVal != nil && pVal != nil {
			sDiff := bst.Difference(sVal.Data, node.Data)
			pDiff := bst.Difference(pVal.Data, node.Data)
			if sDiff <= pDiff {
				kClosest = append(kClosest, sVal)
				i++
			}else{
					kClosest = append(kClosest, pVal)
					j++
			}
		}else if sVal != nil {
			kClosest = append(kClosest, sVal)
			i++
		} else if pVal != nil {
			kClosest = append(kClosest, pVal)
			j++
		}else{
				return kClosest
		}
	}
}

func(bst *BST) kSuccessorsUtil(root, node *BSTNode, k int){
	if root == nil {
		return
	}
	if root == node {
		if root.Right != nil {
			bst.inOrderAddNodesInListUtil(root.Right, k, &kSuccessors)
		}
	} else if bst.Compare(root.Data, node.Data) {
		bst.kSuccessorsUtil(root.Left, node, k)
		kSuccessors = append(kSuccessors, root)
		bst.inOrderAddNodesInListUtil(root.Right, k, &kSuccessors)
	} else {
		bst.kSuccessorsUtil(root.Right, node, k)
	}
}

func(bst *BST) kPredecessorsUtil(root, node *BSTNode, k int){
	if root == nil {
		return
	}
	if root == node {
		if root.Left != nil {
			bst.reverseInOrderAddNodesInListUtil(root.Left, k, &kPredecessors)
		}
	} else if bst.Compare(node.Data, root.Data) {
		bst.kPredecessorsUtil(root.Right, node, k)
		kPredecessors = append(kPredecessors, root)
		bst.reverseInOrderAddNodesInListUtil(root.Left, k, &kPredecessors)
	} else {
		bst.kPredecessorsUtil(root.Left, node, k)
	}
}

func (bst *BST) inOrderAddNodesInListUtil(node * BSTNode, k int , list *[]*BSTNode){
	if len(*list) >= k || node == nil {
		return
	}
	bst.inOrderAddNodesInListUtil(node.Left, k, list)
	if len(*list) >= k {
		return
	}
	*list = append(*list, node)
	bst.inOrderAddNodesInListUtil(node.Right, k, list)
}

func (bst *BST) reverseInOrderAddNodesInListUtil(node * BSTNode, k int , list *[]*BSTNode){
	if len(*list) >= k || node == nil {
		return
	}
	bst.reverseInOrderAddNodesInListUtil(node.Right, k, list)
	if len(*list) >= k {
		return
	}
	*list = append(*list, node)
	bst.reverseInOrderAddNodesInListUtil(node.Left, k, list)
}




