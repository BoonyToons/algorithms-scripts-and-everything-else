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
	getLength()
	isEmpty()
	printQueue()
}

func (queue *Queue) enqueue(val interface{}) {
	newNode := &node{value: val}
	if queue.length == 0 {
		queue.head = newNode 
		queue.tail = newNode
	} else {
		queue.tail.next = newNode
		queue.tail = newNode
	}
	queue.length++
}

func (queue *Queue) dequeue() interface{} {
	var returnValue interface{} = "There are not enough elements in this queue to remove" 
	if queue.length > 0 {
		returnValue = queue.head.value
		queue.head = queue.head.next
		queue.length--
	}

	return returnValue
}

func (queue *Queue) peek() interface{} {
	var returnValue interface{} = "There is no value to return"
	if queue.length != 0 {
		returnValue = queue.head.value
	}

	return returnValue
}

func (queue *Queue) isEmpty() bool {
	returnValue := true

	if queue.length > 0 {
		returnValue = false
	}

	return returnValue
}

func (queue *Queue) getLength() int {
	return queue.length
}
 
func (queue *Queue) printQueue() {
	printValue := queue.head
	if queue.length > 1 {
		for i:=0; i<queue.length; i++ {
			fmt.Printf("%v ", printValue.value)
			printValue = printValue.next
		}
		fmt.Println("")
	} else if queue.length == 1 {
		fmt.Println(printValue.value)
	} else {
		fmt.Println("There are no elements in this queue to print")
	}
}


func main() {
	fmt.Println("---- This is a Queue Data Structure ----")
	newQueue := Queue{}
	fmt.Println(newQueue.dequeue())
	fmt.Println(newQueue.peek())
	newQueue.printQueue()

	fmt.Printf("Is the Queue empty? %v\n", newQueue.isEmpty())
	fmt.Printf("What is the Queue length? %v\n", newQueue.length)

	newQueue.enqueue(15)
	newQueue.enqueue(33)
	newQueue.enqueue("Hello there")
	newQueue.enqueue(true)
	newQueue.printQueue()
	fmt.Printf("What is the Queue length? %v\n", newQueue.length)
	fmt.Println(newQueue.peek())

	newQueue.dequeue()
	fmt.Printf("Is the Queue empty? %v\n", newQueue.isEmpty())

	
	newQueue.printQueue()
	fmt.Printf("What is the Queue length? %v\n", newQueue.length)
}
