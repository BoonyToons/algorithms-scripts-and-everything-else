package main

import (
	"fmt"
)

type Node struct {
	value int
	next *Node
}

func (node *Node) fNode() (int) {
	return node.value
} 

func main() {
	node := Node{15, nil}

	fmt.Print(node.fNode())
}
