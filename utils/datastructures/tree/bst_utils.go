package tree



var kSuccessors []*BSTNode
var kPredecessors []*BSTNode
var inOrderSuccessor *BSTNode
var nodesAtKDistance []*BSTNode

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




func (bst *BST) kSuccessorsUtil(root, node *BSTNode, k int) {
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

func (bst *BST) kPredecessorsUtil(root, node *BSTNode, k int) {
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

func (bst *BST) inOrderAddNodesInListUtil(node *BSTNode, k int, list *[]*BSTNode) {
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

func (bst *BST) reverseInOrderAddNodesInListUtil(node *BSTNode, k int, list *[]*BSTNode) {
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

func (bst *BST) nodesAtKDistUtil(root, node *BSTNode, k int ) int{
	if root == nil {
		return 0
	}
	if root == node {
		addNodesAtKDistToList(root, &nodesAtKDistance, k)
		return 1
	}else if bst.Compare(node.Data , root.Data) == true {
		dist := bst.nodesAtKDistUtil(root.Right, node, k)
		if dist == k {
			nodesAtKDistance = append(nodesAtKDistance, root)
		}else {
			addNodesAtKDistToList(root.Left, &nodesAtKDistance, k-dist-1)
		}
		return dist+1
	} else{
		dist := bst.nodesAtKDistUtil(root.Left, node, k)
		if dist == k {
			nodesAtKDistance = append(nodesAtKDistance, root)
		}else {
			addNodesAtKDistToList(root.Right, &nodesAtKDistance, k-dist-1)
		}
		return dist+1
	}
}

func addNodesAtKDistToList(node *BSTNode, list *[]*BSTNode, k int){
	if node == nil || k<0{
		return
	}
	if k==0{
		*list = append(*list, node)
	}
	addNodesAtKDistToList(node.Left, list, k-1)
	addNodesAtKDistToList(node.Right, list, k-1)
}
