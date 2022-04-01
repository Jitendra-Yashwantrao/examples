package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	//input := []int{5, 6, 7, 9, 8}
	input := []int{3, 2, 4}
	fmt.Println("Results ", twoSum(input, 6))
}

//O(n)
func twoSum(nums []int, target int) []int {

	var answers []int
	numbersMap := make(map[int]int)
	for j := 0; j < len(nums); j++ {
		numbersMap[nums[j]] = j

	}

	for j := 0; j < len(nums); j++ {
		nextMatch := target - nums[j]

		value, present := numbersMap[nextMatch]

		if present && value != j {
			answers = append(answers, j, value)
			break
		}

	}

	return answers
}

func twoSumBruteForce(nums []int, target int) []int {

	var answers []int

	for i, v := range nums {

		for j := i + 1; j < len(nums); j++ {

			if v+nums[j] == target {
				answers = append(answers, i, j)
				return answers
			}
		}
	}

	return answers
}
