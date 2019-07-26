package main

import (
	"fmt"
	"go-google/problems"
	. "go-google/utils/tree"
	"reflect"
)
type TestClass struct {
	val int
}

func (t TestClass) InitNode() *BSTNode {
	return &BSTNode{Data: t}
}


func main(){
	val := TestClass{5}
	bst := InitEmptyBST(reflect.TypeOf(val), func(v1, v2 BSTNodeInterface) bool {
		return v1.(TestClass).val > v2.(TestClass).val
	})
	bst.InsertNode(TestClass{5}.InitNode())
	bst.InsertNode(TestClass{4}.InitNode())
	bst.InsertNode(TestClass{2}.InitNode())
	bst.InsertNode(TestClass{7}.InitNode())
	bst.InsertNode(TestClass{9}.InitNode())
	fmt.Println("==============")
	bst.PrintInOrderTraversal(bst.Root)

	arr := []int{1, 5, 9, 20, 24, 36, 48, 50, 12, 8, 6, 4, 3, 2, 1, 0, -5, -8, -15,-19, -23, -26, -29, -31, -35}

	input := problems.Input{arr, len(arr), 0}
	fmt.Println("Answer : ", input.FindIfExists(-7), ", size of arr = ", len(arr))
	fmt.Println("GetValue method call count : ", input.GetCount)

}
