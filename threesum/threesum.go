package threesum

import (
	"examples/mergesort"
	"fmt"
)

//ThreeSum sum of three nums to zero
func ThreeSum(nums []int) [][]int {
	tripletList := make([][]int, 0, 0)
	sortednums := mergesort.MergeSort(nums)
	fmt.Println("sorted list", sortednums)
	uniquetriplet := make(map[string]string)
	for i := 0; i < len(sortednums); i++ {
		left := i + 1
		right := len(sortednums) - 1

		for left < right {
			sum := sortednums[i] + sortednums[left] + sortednums[right]
			if sum == 0 {

				key := fmt.Sprintf("%v%v%v", sortednums[i], sortednums[left], sortednums[right])
				fmt.Println("triplet found", key)
				_, ispresent := uniquetriplet[key]
				if !ispresent {
					uniquetriplet[key] = key
					tripletList = append(tripletList, []int{sortednums[i], sortednums[left], sortednums[right]})
				}

				left++
			} else if sum < 0 {

				left++
			} else {
				right--
			}

		}
	}

	return tripletList
}
