package main

import (
	"fmt"
)

type node struct {
	value interface{}
	prev *node
}

type Stack struct {
	length int
	head *node
	tail *node
}

type stackInterface interface {
	push()
	pop()
	peek()
	getLength()
	isEmpty()
	printStack()
}

func (stack *Stack) push() {

}

func main() {
	fmt.Println("---- This is a Stack Data Structure ----")
}