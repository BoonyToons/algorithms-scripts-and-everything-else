package main

import (
	"fmt"
)

type node struct {
	value int
	next *node
}

type LinkedList struct {
	length int
	head *node
	tail *node
}

type linkedListInterface interface {
	addLast() *node
	addFirst() *node
	printList() 
}

func (l *LinkedList) addLast(v int) {
	newNode := &node{value: v}
	l.length++
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedList) addFirst(v int) {
	newNode := &node{value: v}
	l.length++
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *LinkedList) printList() {
	printNode := l.head
	for i:=0; i<l.length; i++ {
		fmt.Printf("%v ", printNode.value)
		printNode = printNode.next
	}
	fmt.Println("")
}

func (l *LinkedList) getHead() int {
	return l.head.value
}

func (l *LinkedList) getTail() int {
	return l.tail.value
}

func main() {
	linkedlist := LinkedList{}
	linkedlist.addFirst(15)
	linkedlist.addFirst(6)
	linkedlist.addFirst(22)

	for i:=0; i<5; i++ {
		linkedlist.addFirst(i)
	}
	fmt.Println(linkedlist.getTail())

	linkedlist.addLast(88)
	linkedlist.printList()
	fmt.Println(linkedlist.getTail())
}