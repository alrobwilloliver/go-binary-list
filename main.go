package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 3, 4, 5, 6, 8, 10, 11, 13, 23, 53, 74, 87}
	find := binarySearch(list, 2, len(list)-1, 0)
	fmt.Printf("The target is at index %d\n", find)
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
