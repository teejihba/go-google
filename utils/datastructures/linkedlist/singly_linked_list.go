package linkedlist

import (
	"errors"
	"fmt"
	"reflect"
)

/*
To create a linked list of generic type , create a struct that implements SinglyLinkedListNodeInterface
Checkout example in singly_linked_list.go
*/

type SinglyLinkedListNodeInterface interface {
	InitNode() *SinglyLinkedListNode
	GetData() SinglyLinkedListNodeInterface
}

type SinglyLinkedListNode struct {
	Data SinglyLinkedListNodeInterface
	Next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	Head               *SinglyLinkedListNode
	Tail               *SinglyLinkedListNode
	LinkedListNodeType reflect.Type
}

func (list *SinglyLinkedList) InsertAtEnd(node *SinglyLinkedListNode) error {
	if list.LinkedListNodeType.String() != reflect.TypeOf(node.Data).String() {
		return errors.New("type mismatch")
	}
	if list == nil {
		return errors.New("cannot insert in nil list")
	}
	//empty list
	if list.Tail == nil {
		list.Head = node
		list.Tail = node
		return nil
	}
	list.Tail.Next = node
	list.Tail = node
	return nil
}

func (list *SinglyLinkedList) InsertAtBeginning(node *SinglyLinkedListNode) error {
	if list.LinkedListNodeType.String() != reflect.TypeOf(node.Data).String() {
		return errors.New("type mismatch")
	}
	if list == nil {
		return errors.New("cannot insert in nil list")
	}
	//empty list
	if list.Tail == nil {
		list.Head = node
		list.Tail = node
		return nil
	}
	node.Next = list.Head
	list.Head = node
	return nil
}

func (list *SinglyLinkedList) RemoveNode(node *SinglyLinkedListNode, prevNode *SinglyLinkedListNode) error {
	if list.LinkedListNodeType.String() != reflect.TypeOf(node.Data).String() {
		return errors.New("type mismatch")
	}
	if node == nil {
		return errors.New("cannot remove nil node")
	}
	if list == nil {
		return errors.New("cannot remove in nil list")
	}
	//removing first node
	if prevNode == nil {
		list.Head = node.Next
		//removing the only node
		if node.Next == nil {
			list.Tail = nil
		}
		return nil
	}
	prevNode.Next = node.Next
	node.Next = nil
	return nil
}

func (list *SinglyLinkedList) PrintList() error {
	if list == nil {
		return errors.New("cannot print nil list")
	}
	head := list.Head
	for head != nil {
		fmt.Println("val = ", head.Data)
		head = head.Next
	}
	return nil
}

func (list *SinglyLinkedList) IsListEmpty() bool {
	return list.Head == nil && list.Tail == nil
}

func InitSinglyLinkedList(dataType reflect.Type) *SinglyLinkedList {
	linkedList := SinglyLinkedList{nil, nil, dataType}
	return &linkedList
}
