package main

import (
	"fmt"
)

type node[T any] struct {
	value T
	next *node[T]
}

type Queue[T any] struct {
	length int
	head *node[T]
	tail *node[T]
}

type queueInterface interface {
	enqueue()
	dequeue()
	peek()
}
 

func main() {
	fmt.Println("---- This is a Queue Data Structure ----")

}
