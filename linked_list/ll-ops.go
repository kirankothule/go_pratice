package main

import "fmt"

func (ll *LinkedList) findNthElementFromLast(n int) *node {
	len := 0
	for current := ll.head; current != nil; current = current.next {
		len++
	}
	if n > len {
		return nil
	}
	current := ll.head
	for i := 0; i < len-n; i++ {
		current = current.next
	}
	return current
}

func (ll *LinkedList) findCycleInLinkedList() *node {
	fastPtr := ll.head
	slowPtr := ll.head

	for slowPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if fastPtr == slowPtr {
			return fastPtr
		}
	}
	return nil
}
func (ll *LinkedList) findStartOfLoop() *node {

	if fastPtr := ll.findCycleInLinkedList(); fastPtr != nil {
		slowPrt := ll.head
		for slowPrt != fastPtr {
			slowPrt = slowPrt.next
			fastPtr = fastPtr.next
		}
		return slowPrt
	}
	return nil
}

func (ll *LinkedList) reverseLinkedList() {
	current := ll.head
	var prev *node
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	ll.head = prev
}

func (ll *LinkedList) reverseRecusrssion() {
	ll.head = reverse(ll.head)
}
func reverse(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	rest := reverse(head.next)
	head.next.next = head
	head.next = nil
	return rest
}

func findIntersectingPoint(list1, list2 *LinkedList) *node {
	len1 := list1.length()
	len2 := list2.length()
	head1 := list1.head
	head2 := list2.head
	diff := len1 - len2
	if len1 < len2 {
		head1 = list2.head
		head2 = list1.head
		diff = len2 - len1
	}
	for i := 0; i < diff; i++ {
		head1 = head1.next
	}
	for head1 != nil && head2 != nil {
		if head1 == head2 {
			break
		}
		head1 = head1.next
		head2 = head2.next
	}
	return head1
}

func (ll *LinkedList) getMiddleOfList() *node {
	if ll.head == nil {
		return nil
	}
	if ll.head.next == nil {
		return ll.head
	}
	slowPtr := ll.head
	fastPtr := ll.head

	for fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
	}
	return slowPtr
}

func printReverse(head *node) {
	if head == nil {
		return
	}
	printReverse(head.next)
	fmt.Println(head.data)
}

func mergeSortedList(head1, head2 *node) *node {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	head := &node{
		data: 0,
		next: nil,
	}
	fmt.Println("head1:", head1.data)
	fmt.Println("head2:", head2.data)
	if head1.data <= head2.data {
		head = head1
		head.next = mergeSortedList(head1.next, head2)
	} else {
		head = head2
		head.next = mergeSortedList(head2.next, head1)
	}
	return head
}

// func (ll *LinkedList) reverseKElements(k int) {

// 	kList := ll.head
// 	restList := ll.head

// 	for i := 0; i < k; i++ {
// 		restList = restList.next
// 	}
// 	temp := restList
// 	restList = restList.next
// 	temp.next = nil
// 	reverseHead := reverse(kList)
// }
