package main

import (
	"errors"
	"fmt"
)

type ArrayList[arrType any] struct {
	itemLength int
	arrayLength int
	newArray []arrType
}

func NewArrayList[arrType any](arrayLength int) (*ArrayList[arrType], error) {
	
	if arrayLength < 1 {
		return nil, errors.New("The initial array size is too small")
	}

	return &ArrayList[arrType] {
		itemLength: 0,
		arrayLength: arrayLength,
		newArray: make([]arrType, arrayLength),
	},
	nil
}

type arrayListInterface interface {
	push()
}



func main() {
	fmt.Println("---- This is an ArrayList Data Structure ----")

	list, err := NewArrayList[int](0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(list)
	}
}