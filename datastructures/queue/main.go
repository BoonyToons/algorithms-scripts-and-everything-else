package main

import (
	"fmt"
)

type node struct {
	value interface{}
	next *node
}

type Queue struct {
	length int
	head *node
	tail *node
}

type queueInterface interface {
	enqueue()
	dequeue()
	peek()
	printQueue()
}

func (queue *Queue) enqueue(val interface{}) {
	newNode := &node{value: val}
	if queue.length == 0 {
		queue.head = newNode 
		queue.tail = newNode
	} else {
		newNode.next = queue.tail
		queue.tail = newNode
	}
	queue.length++
}

func (queue *Queue) dequeue() interface{} {
	returnValue := "There are not enough elements in this queue to remove"
	if queue.length > 0 {
		returnValue = queue.head.value
		
		queue.length--
	}
	return returnValue
}

func (queue *Queue) peek() interface{} {
	return queue.head.value
}
 
func (queue *Queue) printQueue() {
	printValue := queue.tail
	if queue.length > 1 {
		for i:=0; i<queue.length; i++ {
			fmt.Printf("%v ", printValue.value)
			printValue = printValue.next
		}
		fmt.Println("")
	} else if queue.length == 1 {
		fmt.Println("%v ", printValue.value)
	} else {
		fmt.Println("There are no elements in this queue to print")
	}
	
}


func main() {
	fmt.Println("---- This is a Queue Data Structure ----")
	newQueue := Queue{}
	newQueue.enqueue(15)
	newQueue.enqueue(33)
	newQueue.enqueue("Hello there")
	newQueue.enqueue(true)

	fmt.Println(newQueue.peek())
	newQueue.printQueue()
}
