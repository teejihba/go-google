package main
//
//import (
//	"errors"
//	"fmt"
//	"reflect"
//	"strings"
//)
//
//type MyNode struct {
//	value int
//}
//
//func (node *MyNode) PrintValue() {
//	fmt.Printf(" %d ", node.value)
//}
//
//type llNode struct {
//	key        llNodeInterface
//	next       *llNode
//	llNodeType reflect.Type
//}
//
//type llNodeInterface interface {
//	PrintValue()
//}
//
//type ComplexNode struct {
//	realValue  int
//	imageValue int
//}
//
//func (node *ComplexNode) PrintValue() {
//	fmt.Printf(" %d + i%d", node.realValue, node.imageValue)
//}
//
//// Student type.
//type Student struct {
//	name string
//	age  int
//}
//
//// Student implements the PrintValue function - thus llNodeInterface is implemented.
//func (node *Student) PrintValue() {
//	fmt.Printf("Name: %s | Age : %d ", node.name, node.age)
//}
//
//// Function which will check the of the new node before adding to the linked
//// list. It checks the type of the new key against the type of the key in the
//// head. If both are equal then it proceed else return error.
//func (head *llNode) AddBeforeHeadTypeCheck(passedKey llNodeInterface) error {
//
//	if head.key == nil {
//		head.key = passedKey
//		head.llNodeType = reflect.TypeOf(head.key)
//	} else {
//		typeOfPassedKey := reflect.TypeOf(passedKey)
//
//		if typeOfPassedKey != head.llNodeType {
//			fmt.Printf("\nUnsupported type for the type %T", passedKey)
//			return errors.New("Type mistmatch")
//		}
//
//		temp := llNode{key: head.key, next: head.next}
//		head.key = passedKey
//		head.next = &temp
//	}
//	return nil
//}
//
//// Function which will not check the types and will simply add the new node to
//// the linked list. Thus linked list will be able to have nodes of multiple
//// types.
//func (head *llNode) AddBeforeHead(passedKey llNodeInterface) {
//
//	if head.key == nil {
//		head.key = passedKey
//		head.llNodeType = reflect.TypeOf(head.key)
//	} else {
//		temp := llNode{key: head.key, next: head.next}
//		head.key = passedKey
//		head.next = &temp
//	}
//}
//
//func (head *llNode) Init() {
//	head.key = nil
//	head.next = nil
//	head.llNodeType = nil
//
//}
//
//// Print the linked list.
//func (head *llNode) DisplayLL() {
//
//	temp := head
//	fmt.Printf("\n%s", strings.Repeat("#", 80))
//	fmt.Printf("\nPrinting the linked list\n")
//
//	for {
//		if temp.key == nil {
//			fmt.Println("Linked list is empty")
//			break
//		} else {
//			fmt.Printf("\n %T %v ", temp.key, temp.key)
//			key := temp.key
//			key.PrintValue()
//			if temp.next == nil {
//				break
//			} else {
//				temp = temp.next
//			}
//		}
//	}
//	fmt.Printf("\n%s", strings.Repeat("#", 80))
//	fmt.Printf("\n\n")
//}
//
//func testWithMixedType() {
//	head := llNode{}
//	head.Init()
//
//	for i := 1; i < 10; i++ {
//		temp := &ComplexNode{i, i * 10}
//		head.AddBeforeHeadTypeCheck(temp)
//	}
//
//	temps := &Student{"rishi", 20}
//	head.AddBeforeHeadTypeCheck(temps) // Will give error.
//	head.DisplayLL()
//}
//
//func testWithComplexNumber() {
//
//	head := llNode{}
//	head.Init()
//
//	for i := 1; i < 10; i++ {
//		temp := &ComplexNode{i, i * 10}
//		head.AddBeforeHeadTypeCheck(temp)
//	}
//
//}
//
//func main() {
//	testWithComplexNumber()
//	testWithMixedType()
//
//}
