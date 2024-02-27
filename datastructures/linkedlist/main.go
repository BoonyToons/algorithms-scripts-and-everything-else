package main

import (
	//"fmt"
)

type node struct {
	value int
	next *node
}


type LinkedList struct {
	length int
	head *node
}


type linkedListInterface interface {
	addLast() *node
}

func (l *LinkedList) addLast(v int) {
	node := &node{value: v}
}


func main() {
	linkedlist := LinkedList{}
	linkedlist.addLast(15)
}