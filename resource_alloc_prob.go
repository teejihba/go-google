package main

import (
	"fmt"
	"go-google/utils"
	"reflect"
	"sort"
)

type Process struct {
	Name string
	Allocated int
	Required int
}

func (p Process) InitNode() *utils.SinglyLinkedListNode {
	node := &utils.SinglyLinkedListNode{Data: p, Next:nil}
	return node
}

func (p Process) GetData() utils.LinkedListNodeInterface {
	return p
}

func main(){
	//////problem input////////
	n := 4
	processes := [] string {"p1","p2", "p3","p4"}
	allocatedArr := [] int {1,2,3,4}
	requiredArr := [] int {5,2,1,5}
	available := 1
	k := 2
	///////////////////////////
	list := sortProcArrayAndInitLinkedList(n, processes, allocatedArr, requiredArr)
	curr := list.Head
	var prev *utils.SinglyLinkedListNode = nil
	for !list.IsListEmpty() &&  k > 0 && curr != nil {
		var p = curr.Data.GetData().(Process)
		if p.Required <= available{
			available += p.Allocated
			_ =list.RemoveNode(curr, prev)
			k--
			curr = list.Head
			prev = nil
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	if curr == nil && k>0 {
		fmt.Println("Cannot complete k process")
	} else {
		fmt.Println("Answer = ", available)
	}

}

func sortProcArrayAndInitLinkedList(n int, procs []string, alloc []int, reqd []int) *utils.SinglyLinkedList {
	procArr := initProcArr(n, procs, alloc, reqd)
	sort.SliceStable(procArr, func(i, j int) bool {
		return procArr[i].Allocated > procArr[j].Allocated
	})
	list := utils.InitSinglyLinkedList(reflect.TypeOf(Process{}))
	for i := 0; i < n; i++ {
		node := procArr[i].InitNode()
		_ = list.InsertAtEnd(node)
	}
	return list
}

func initProcArr(n int, procs []string, alloc []int, reqd []int) [] Process {
	procArr := make([]Process, n)
	for i := 0; i < n; i++ {
		procArr[i] = Process{procs[i], alloc[i], reqd[i]}
	}
	return procArr
}

