package main

import "fmt"

func main() {
	var a = [6]int {6,5,4,3,2,1}
	var b = [6]int {2,5,1,2,1,6}
	var c = [6]int {3,3,2,1,3,3}
	var d = [6]int {1,1,4,4,2,2}
	var e = [6]int {2,1,4,3,6,5}
	bubbleSort(a)
	bubbleSort(b)
	bubbleSort(c)
	bubbleSort(d)
	bubbleSort(e)
}

func bubbleSort(haystack [6]int) {
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