package binarysearch

import "fmt"

func BinarySearchItem(item int, itemList []int) bool {

	low := 0
	high := len(itemList) - 1

	for low <= high {

		median := (low + high) / 2
		if itemList[median] < item {
			low = median + 1
		} else {
			high = median - 1
		}

	}

	if low == len(itemList) || itemList[low] != item {
		return false
	}
	fmt.Println("item found at index", low)
	return true

}
