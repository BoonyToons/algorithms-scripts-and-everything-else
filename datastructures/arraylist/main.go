package main

import (
	"fmt"
)

type ArrayList[T any] struct {
	itemLength int
	arrayLength int
	newArray []T
}

func NewArrayList[T any](arrayLength int) (*ArrayList[T], error) {
	returnValue := &ArrayList[T] {
		itemLength: 0,
		arrayLength: arrayLength,
		newArray: make([]T, arrayLength),
	}

	var err error

	if arrayLength < 1 {
		returnValue = nil
		err = "Hello"
	}

	return returnValue, err
}

type arrayListInterface interface {
	add()
	insert()
	get()
	set()
	remove()
	contains()
	indexOf()
	length()
	isEmpty()
	clear()
}


func (list *ArrayList[T]) add() {
	fmt.Println("Hello")
}


func main() {
	fmt.Println("---- This is an ArrayList Data Structure ----")

	list, err := NewArrayList[int](6)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	list.add()
}