package problems

type Node struct {
	Val int
	Next* Node
}


func IsLinkedListPalindrome(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	return recUtil_(head, &head)
}

func recUtil_(head *Node, cur **Node) bool {
	if head == nil {
		return true
	}
	ret := recUtil_(head.Next, cur)
	if ret == true{
		if (*(*cur)).Val == head.Val {
			*cur = (*cur).Next
			return true
		}
		return false
	}else {
		return false
	}

}

