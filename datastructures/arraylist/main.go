package main

import (
	"fmt"
)

type ArrayList[T any] struct {
	itemLength int
	arrayLength int
	newArray []T
}

type arrayListInterface interface {
	NewArrayList()
}

func (list *ArrayList) NewArrayList(initialSize int) *ArrayList {
	return &list{
		itemLength: 0
		arrayLength: initialSize
		newArray: []
	}
}



func main() {
	fmt.Println("---- This is an ArrayList Data Structure ----")
}