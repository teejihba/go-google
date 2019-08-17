package problems




  //Definition for a binary tree node.
  type TreeNodeTemp struct {
      Val int
      Left *TreeNodeTemp
      Right *TreeNodeTemp
  }

func bstToGst(root *TreeNodeTemp) *TreeNodeTemp {
	sum:= 0
	util(root, &sum)
	return root
}

func util(root *TreeNodeTemp, sum *int){
	if root == nil{
		return
	}
	util(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	util(root.Left, sum)
}
