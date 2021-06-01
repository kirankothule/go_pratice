package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
}

func (list *LinkedList) add(data int) {
	temp := &LinkedList{
		&node{
			data: data,
			next: nil,
		},
	}
	if list.head == nil {
		list.head = &node{
			data: data,
			next: nil,
		}
	} else {

		current := list.head
		for ; current.next != nil; current = current.next {
		}
		current.next = temp.head
	}
}

func (list *LinkedList) display() {
	fmt.Println()
	cnt := 0
	for current := list.head; current != nil; current = current.next {
		fmt.Println(current.data)
		if cnt > 100 {
			break
		}
		cnt++
	}
	fmt.Println()
}

func (list *LinkedList) addAt(data, position int) {
	temp := &node{
		data: data,
		next: nil,
	}
	if position == 0 {
		newNode := &node{
			data: data,
			next: nil,
		}
		newNode.next = list.head.next
		list.head = newNode
	} else {

		current := list.head
		i := 1
		for ; current.next != nil && i < position-1; i, current = i+1, current.next {
		}
		if current == nil {
			return
		}
		temp.next = current.next
		current.next = temp
	}
}

func (list *LinkedList) delete(position int) {
	current := list.head
	if current == nil {
		return
	}
	for i := 1; i < position-1 && current != nil; i, current = i+1, current.next {
	}
	if current == nil {
		return
	}
	nodeToDelete := current.next
	current.next = nodeToDelete.next
	nodeToDelete = nil

}

func (list *LinkedList) getLastElement() *node {
	current := list.head
	for ; current.next != nil; current = current.next {
	}
	return current
}

func (list *LinkedList) getNthElement(position int) *node {

	if position == 0 {
		return list.head
	}

	current := list.head
	i := 1
	for ; current.next != nil && i < position; i, current = i+1, current.next {
	}
	if current == nil {
		return nil
	}
	return current

}

func (list *LinkedList) length() int {

	current := list.head
	i := 1
	for ; current != nil; i, current = i+1, current.next {
	}
	return i
}
