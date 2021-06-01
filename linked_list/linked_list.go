package main

import "fmt"

func main() {
	list1 := &LinkedList{}
	list1.add(10)
	list1.add(20)
	list1.add(30)
	list1.add(40)
	list1.add(50)
	list1.add(60)
	list1.add(70)
	fmt.Println("printing")
	list1.display()

	list2 := &LinkedList{}
	list2.add(11)
	list2.add(21)
	list2.add(31)
	list2.add(41)
	list2.add(51)
	list2.add(61)
	list2.add(71)
	fmt.Println("printing")
	list2.display()

	// // Merge list
	// list3 := &LinkedList{
	// 	head: mergeSortedList(list1.head, list2.head),
	// }
	// fmt.Println("printing mearged list")
	// list3.display()

	// fmt.Println("Print in reverse")
	// printReverse(list1.head)

	// // Middle of list
	// ele := list1.getMiddleOfList()
	// fmt.Println("The middle of list list: ", ele)

	// // intersection point
	// list2.getLastElement().next = list1.getNthElement(3)
	// fmt.Println("printing")
	// list2.display()

	// ele := findIntersectingPoint(list1, list2)
	// if ele == nil {
	// 	fmt.Println("There is no intersecting point")
	// } else {
	// 	fmt.Println("The intersecting point is: ", ele.data)
	// }
	// //end

	// list.reverseRecusrssion()
	// fmt.Println("reverse printing")
	// list.display()

	// list.reverseLinkedList()
	// fmt.Println("reverse printing")
	// list.display()

	// ele := list.getNthElement(3)

	// fmt.Println("the nth element is: ", ele.data)
	// list.getLastElement().next = ele

	// ele = list.findCycleInLinkedList()
	// if ele == nil {
	// 	fmt.Println("Linked list does not have cycle")
	// } else {
	// 	fmt.Println("Linked list have cycle:", ele.data)
	// }

	// ele = list.findStartOfLoop()
	// if ele == nil {
	// 	fmt.Println("loop not present in list")
	// } else {
	// 	fmt.Println("Linked list have cycle and start point is:", ele.data)
	// }
	// ele := list.findNthElementFromLast(5)
	// fmt.Println(ele.data)

	// Stack using linked list
	// var s stack
	// s.push(10)
	// s.push(20)
	// s.push(30)
	// s.display()
	// s.pop()
	// s.display()
	// s.pop()
	// s.display()
	// s.pop()
	// s.display()
	// s.pop()
	// s.display()
	// s.pop()
	// s.display()
}
