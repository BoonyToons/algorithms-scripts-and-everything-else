package main

import (
	"fmt"
)

type ArrayList[T any] struct {
	itemLength int
	arrayLength int
	newArray []T
}

func NewArrayList[T any](arrayLength int) *ArrayList[T] {
	return &ArrayList[T] {
		itemLength: 0,
		arrayLength: arrayLength,
		newArray: make([]T, arrayLength),
	}
}

type arrayListInterface interface {
	push()
}



func main() {
	fmt.Println("---- This is an ArrayList Data Structure ----")

	list := NewArrayList[int](6)
	fmt.Println(list)
}