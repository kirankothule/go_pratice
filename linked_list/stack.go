package main

import "fmt"

type stack struct {
	top *node
}

func (s *stack) iSEmpty() bool {
	return s.top == nil
}

func (s *stack) push(data int) bool {
	temp := s.top
	newNode := &node{
		data: data,
		next: nil,
	}
	newNode.next = temp
	s.top = newNode
	return true
}
func (s *stack) pop() int {
	if s.iSEmpty() {
		fmt.Println("foubd empty")
		return 0
	}
	fmt.Println("here to pop element")
	podEle := s.top
	data := podEle.data
	s.top = podEle.next
	podEle = nil
	return data
}
func (s *stack) display() {
	temp := s.top
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
