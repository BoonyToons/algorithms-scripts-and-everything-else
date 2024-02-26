package main

import (
	"fmt"
)

type Node struct {
	value int
	next struct
}

func (node *Node) fNode() (int) {
	return node.value
} 

func main() {
	node := Node{15}

	fmt.Print(node.fNode())
}