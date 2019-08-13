package main

import (
	"fmt"
	"go-google/problems"
	. "go-google/utils/datastructures/tree"
)

type TestClass struct {
	val int
}

func (t TestClass) InitNode() *BSTNode {
	return &BSTNode{Data: t}
}

func main() {
	//val := TestClass{5}
	//cmp := func(v1, v2 BSTNodeInterface) bool {
	//	return v1.(TestClass).val > v2.(TestClass).val
	//}
	//diff := func(v1, v2 BSTNodeInterface) int {
	//	i := v1.(TestClass).val
	//	j := v2.(TestClass).val
	//	diff := i - j
	//	return mathsUtil.AbsoluteVal(diff)
	//}
	//bst := InitEmptyBST(reflect.TypeOf(val), cmp, diff)
	//bst.InsertNode(TestClass{5}.InitNode())
	//bst.InsertNode(TestClass{4}.InitNode())
	//node2 := TestClass{2}.InitNode()
	//bst.InsertNode(node2)
	//node7 := TestClass{7}.InitNode()
	//bst.InsertNode(node7)
	//node9 := TestClass{9}.InitNode()
	//bst.InsertNode(node9)
	//node19 := TestClass{19}.InitNode()
	//bst.InsertNode(node19)
	//bst.InsertNode(TestClass{13}.InitNode())
	//bst.InsertNode(TestClass{8}.InitNode())
	//node99 := TestClass{99}.InitNode()
	//bst.InsertNode(node99)
	//bst.InsertNode(TestClass{1}.InitNode())
	//node44 := TestClass{44}.InitNode()
	//bst.InsertNode(node44)
	//bst.PrintInOrderTraversal(bst.Root)
	//fmt.Println("==============")

	//fmt.Println(bst.InOrderSuccessor(bst.Root, node9))
	//fmt.Println(bst.InOrderSuccessor(bst.Root, node19))
	//fmt.Println(bst.InOrderSuccessor(bst.Root, node2))
	//fmt.Println(bst.InOrderSuccessor(bst.Root, node44))
	//fmt.Println(bst.InOrderSuccessor(bst.Root, node99))

	//successors := bst.KInOrderSuccessors(bst.Root, node9, 4)
	//for i := range successors {
	//	fmt.Println("----", *successors[i])
	//}
	//
	//predecessors := bst.KInOrderPredecessors(bst.Root, node9, 4)
	//for i := range predecessors {
	//	fmt.Println("****", *predecessors[i])
	//}
	//
	//kClosest := bst.KClosestNodeOfANode(bst.Root, node9, 4)
	//for i := range kClosest {
	//	fmt.Println("+++++++", *kClosest[i])
	//}
	//
	//nodesAtK := bst.NodesAtKDistance(bst.Root, node44, 2)
	//for i:= range nodesAtK{
	//	fmt.Println("$$$$$", *nodesAtK[i])
	//}
	//arr := []int{1, 5, 9, 20, 24, 36, 48, 50, 12, 8, 6, 4, 3, 2, 1, 0, -5, -8, -15,-19, -23, -26, -29, -31, -35}
	//
	//input := problems.Input{arr, len(arr), 0}
	//fmt.Println("Answer : ", input.FindIfExists(-7), ", size of arr = ", len(arr))
	//fmt.Println("GetValue method call count : ", input.GetCount)
	//i1, e := mathsUtil.InitBigInt("90")
	//if e != nil{
	//	fmt.Println("err = ", e)
	//}
	//i2, _ := mathsUtil.InitBigInt("4")
	//sum, _ := i1.Add(i2)
	//fmt.Println(i1,i2,sum)
	//words := []string{"area","lead","wall","lady","ball"}
	//str := "adsfadasd"
	//
	//i, i2 := problems.LongestSubstringWithAtmostKDisntinctChar(str, 3)
	//fmt.Println("", i, str[i2:i2+i])
	//[]int{9,10,11,12},[]int{13,14,15,16}
	//grid := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8},{9,10,11,12}}
	//tree := quadtree.InitQuadTree(grid, 3, 4)
	//fmt.Println("Tree == ",  tree.Tree)
	//
	//val  := tree.RangeQuery(0,1,1,2)
	//fmt.Println("Val == ",  val)
	//
	//val  = tree.RangeQuery(0,1,2,3)
	//fmt.Println("Val == ",  val)
	//tree.PointUpdate(0,0, 7)
	//fmt.Println("New tree = ", tree.Tree)
	//
	//MySet := disjointset.DisjointSet{}
	//MySet.MakeSet(6)
	//fmt.Println(MySet)
	//
	//MySet.Union(3,5)
	//fmt.Println(MySet)
	//MySet.Union(1,2)
	//fmt.Println(MySet)
	//MySet.Union(2,3)
	//fmt.Println(MySet)
	//
	//fmt.Println("1 = %T 5 = ? 3 = ?",MySet.Find(1),MySet.Find(5),MySet.Find(3))

	yes, st, end := problems.KmpSubstrSearch("acaacabac","acabac")
	fmt.Println(yes, st, end)
	////arr := []int{4,5,9,10,15,17,19,34,98}
	//var arr []TestClass
	//classes := []TestClass{{34}, {15}, {10}, {9}, {5}, {4}}
	//arr = append(arr, classes[0],classes[1],classes[2],classes[3],classes[4], classes[5])
	//
	//ind := sort.Search(len(arr), func(i int) bool {
	//	return arr[i].val <= 8
	//})
	//fmt.Println(arr[ind:])
	//arr = append(arr, TestClass{})
	//copy(arr[ind+1:], arr[ind:])
	//fmt.Println(arr)
	//arr[ind] = TestClass{8}
	//fmt.Println(arr)
	//fmt.Println("indesx = ", ind)
	//fmt.Println(strings.HasPrefix())
}
