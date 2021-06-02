package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fs := flag.NewFlagSet("binary-list", flag.ExitOnError)
	var i int
	fs.IntVar(&i, "t", 1, "The target to look for")
	fs.PrintDefaults()
	fs.Parse(os.Args[1:])
	list := []int{1, 2, 3, 3, 4, 5, 6, 8, 10, 11, 13, 23, 53, 74, 87}
	// if !Contains(list, i) {
	// 	fmt.Println("The target doesn't appear in the list!")
	// 	return
	// }
	find := binarySearch(list, i, len(list)-1, 0)
	fmt.Printf("The target is at index %d\n", find)
}

func Contains(list []int, x int) bool {
	for _, item := range list {
		if item == x {
			return true
		}
	}
	return false
}

func binarySearch(array []int, target int, highIndex int, lowIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := int((highIndex + lowIndex) / 2)
	if array[mid] < target {
		return binarySearch(array, target, highIndex, mid+1)
	} else if array[mid] > target {
		return binarySearch(array, target, mid, lowIndex)
	} else {
		return mid
	}
}
