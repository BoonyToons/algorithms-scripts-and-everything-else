package main

import (
	"fmt"
)

type ArrayList[arrType any] struct {
	itemLength int
	arrayLength int
	newArray []arrType
}

func NewArrayList[arrType any](arrayLength int) *ArrayList[arrType] {
	return &ArrayList[arrType] {
		itemLength: 0,
		arrayLength: arrayLength,
		newArray: make([]arrType, arrayLength),
	}
}

type arrayListInterface interface {
	push()
}



func main() {
	fmt.Println("---- This is an ArrayList Data Structure ----")

	list := NewArrayList[int](10)
	fmt.Println(list)
}