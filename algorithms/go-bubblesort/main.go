package main

import (
	"fmt"
)

func main() {
	var array1 = [6]int{6,5,4,3,2,1}
	bubbleSort(array1)
}


func bubbleSort(list [6]int) [6]int {

	var tempList [6]int = list
	for i:=0; i<len(tempList); i++ {
		for j:=0; j<len(tempList) - (1-i); j++ {
			if tempList[j] > tempList[j+1] {
				tempList[j], tempList[j+1] := intSwap(tempList[j], tempList[j+1])
			}
		}
	}

	for _, v := range tempList {
		fmt.Printf("%i", v)
		fmt.Print(" ")
	}
	fmt.Println()

	return tempList
}

func intSwap(a, b int) (int, int) {
	return b, a
}