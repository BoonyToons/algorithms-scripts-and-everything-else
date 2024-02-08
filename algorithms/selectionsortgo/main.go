package main

import "fmt"

func main() {
	var a = [6]int {6,5,4,3,2,1}
	var b = [6]int {2,5,1,2,1,6}
	var c = [6]int {3,3,2,1,3,3}
	var d = [6]int {1,1,4,4,2,2}
	var e = [6]int {2,1,4,3,6,5}
	selectionSort(a)
	selectionSort(b)
	selectionSort(c)
	selectionSort(d)
	selectionSort(e)
}

func selectionSort(haystack [6]int) {
	for ii:=0; ii<6; ii++ {
		var low = haystack[ii]
		var index = ii
		for jj:=ii; jj<6 - 1; jj++ {
			if low > haystack[jj+1] {
				low = haystack[jj+1]
				index = jj+1
			}
		}
		haystack[index] = haystack[ii]
		haystack[ii] = low
	}

	for ii:=0; ii<6; ii++ {
		fmt.Print(haystack[ii])
		fmt.Print(" ")
	}
	fmt.Println("")
}