package main

import (
	"fmt"
	"go-google/utils/datastructures/disjointset"
	quadtree "go-google/utils/datastructures/quad-tree"
	. "go-google/utils/datastructures/tree"
)
type TestClass struct {
	val int
}

func (t TestClass) InitNode() *BSTNode {
	return &BSTNode{Data: t}
}


func main(){
	//val := TestClass{5}
	//bst := InitEmptyBST(reflect.TypeOf(val), func(v1, v2 BSTNodeInterface) bool {
	//	return v1.(TestClass).val > v2.(TestClass).val
	//})
	//bst.InsertNode(TestClass{5}.InitNode())
	//bst.InsertNode(TestClass{4}.InitNode())
	//bst.InsertNode(TestClass{2}.InitNode())
	//bst.InsertNode(TestClass{7}.InitNode())
	//bst.InsertNode(TestClass{9}.InitNode())
	//fmt.Println("==============")
	//bst.PrintInOrderTraversal(bst.Root)
	//
	//arr := []int{1, 5, 9, 20, 24, 36, 48, 50, 12, 8, 6, 4, 3, 2, 1, 0, -5, -8, -15,-19, -23, -26, -29, -31, -35}
	//
	//input := problems.Input{arr, len(arr), 0}
	//fmt.Println("Answer : ", input.FindIfExists(-7), ", size of arr = ", len(arr))
	//fmt.Println("GetValue method call count : ", input.GetCount)
	//i1, e := math.InitBigInt("90")
	//if e != nil{
	//	fmt.Println("err = ", e)
	//}
	//i2, _ := math.InitBigInt("4")
	//sum, _ := i1.Add(i2)
	//fmt.Println(i1,i2,sum)
	//words := []string{"area","lead","wall","lady","ball"}
	//str := "adsfadasd"
	//
	//i, i2 := problems.LongestSubstringWithAtmostKDisntinctChar(str, 3)
	//fmt.Println("", i, str[i2:i2+i])
	//[]int{9,10,11,12},[]int{13,14,15,16}
	grid := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8},{9,10,11,12}}
	tree := quadtree.InitQuadTree(grid, 3, 4)
	fmt.Println("Tree == ",  tree.Tree)

	val  := tree.RangeQuery(0,1,1,2)
	fmt.Println("Val == ",  val)

	val  = tree.RangeQuery(0,1,2,3)
	fmt.Println("Val == ",  val)
	tree.PointUpdate(0,0, 7)
	fmt.Println("New tree = ", tree.Tree)

	MySet := disjointset.DisjointSet{}
	MySet.MakeSet(6)
	fmt.Println(MySet)

	MySet.Union(3,5)
	fmt.Println(MySet)
	MySet.Union(1,2)
	fmt.Println(MySet)
	MySet.Union(2,3)
	fmt.Println(MySet)

	fmt.Println("1 = %T 5 = ? 3 = ?",MySet.Find(1),MySet.Find(5),MySet.Find(3))



}
