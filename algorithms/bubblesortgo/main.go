package main

import "fmt"

func main() {
	var a = []int {6,5,4,3,2,1}
	var b = []int {2,5,1,2,1,6}
	var c = []int {3,3,2,1,3,3}
	var d = []int {1,1,4,4,2,2}
	var e = []int {2,1,4,3,6,5}
	bubbleSort(a)
	bubbleSort(b)
	bubbleSort(c)
	bubbleSort(d)
	bubbleSort(e)
}

func bubbleSort(haystack []int) {
	for ii := 0; ii<6; ii++ {
		for jj := 0; jj<6 - 1 - ii; jj++ {
			if haystack[jj] > haystack[jj+1] {
				var left = haystack[jj];
				var right = haystack[jj+1];
				haystack[jj] = right;
				haystack[jj+1] = left;
			}
		}
	}

	for ii:=0; ii<6; ii++ {
		fmt.Print(haystack[ii])
		fmt.Print(" ")
	}
	fmt.Println(" ")
}