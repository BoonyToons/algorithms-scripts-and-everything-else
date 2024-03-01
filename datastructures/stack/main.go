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

func (stack *Stack) push(val interface{}) interface{} {
	newNode := &node{value: val}
	if stack.length > 0 {
		newNode.prev = stack.head
		stack.head = newNode

	} else {
		stack.head = newNode
		stack.tail = newNode
	}

	stack.length++

	return newNode.value
}

func (stack *Stack) pop() interface{} {
	var returnValue interface{} = "There is not enough elements in this stack to pop"

	if stack.length > 0 {
		returnValue = stack.head.value
		stack.head = stack.head.prev
		stack.length--
	} 

	return returnValue
}

func (stack *Stack) peek() interface{} {
	var returnValue interface{} = "There is not enough elements in the Stack to peek"

	if stack.length > 0 {
		returnValue = stack.head.value
	} 

	return returnValue
}

func (stack *Stack) getLength() int {
	return stack.length
}

func (stack *Stack) isEmpty() bool {
	var returnValue bool = true

	if stack.length > 0 {
		returnValue = false
	}

	return returnValue
}

func (stack *Stack) printStack() {
	printValue := stack.head
	if stack.length > 1 {
		for i:=0; i<stack.length; i++ {
			fmt.Printf("%v ", printValue.value)
			printValue = printValue.prev
		}
		fmt.Println("")
	} else if stack.length == 1 {
		fmt.Println(printValue.value)
	} else {
		fmt.Println("There are no elements in this stack to print")
	}
}

func main() {
	fmt.Println("---- This is a Stack Data Structure ----")

	newStack := Stack{}

	newStack.printStack()
	fmt.Printf("'%v' is the length of the stack\n", newStack.getLength())
	fmt.Printf("Is the stack empty? %v\n", newStack.isEmpty())
	fmt.Printf("'%v' was added to the Stack\n", newStack.push(5))
	fmt.Printf("'%v' was added to the Stack\n", newStack.push("hello world"))
	fmt.Printf("'%v' was added to the Stack\n", newStack.push(true))
	newStack.printStack()
	fmt.Printf("'%v' is the length of the stack\n", newStack.getLength())
	fmt.Printf("Is the stack empty? %v\n", newStack.isEmpty())

	fmt.Printf("'%v' is the value on top of the stack\n", newStack.peek())

	fmt.Printf("'%v' has been popped!\n", newStack.pop())
	fmt.Printf("'%v' has been popped!\n", newStack.pop())
	newStack.printStack()
	fmt.Printf("'%v' is the length of the stack\n", newStack.getLength())
	fmt.Printf("Is the stack empty? %v\n", newStack.isEmpty())
	fmt.Printf("'%v' is the value on top of the stack\n", newStack.peek())
}