package main

import (
	"testing"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
*/

func TestTwoSum(t *testing.T) {

	input := []int{3, 2, 4}
	expectedOutput := []int{1, 2}
	output := twoSum(input, 6)
	//fmt.Println(output)
	t.Log(output)
	if len(output) != len(expectedOutput) {
		t.Fail()
	}

	if (output[0] != expectedOutput[0]) || (output[1] != expectedOutput[1]) {
		t.Fail()
	}

}

func TestTwoSum_duplicateNumber(t *testing.T) {

	input := []int{3, 3}
	expectedOutput := []int{0, 1}
	output := twoSum(input, 6)
	//fmt.Println(output)
	t.Log(output)
	if len(output) != len(expectedOutput) {
		t.Fail()
	}
	// if (output[0] != expectedOutput[0]) || (output[1] != expectedOutput[1]) {
	// 	t.Fail()
	// }
}

func TestTwoSum_duplicateNumber2(t *testing.T) {

	input := []int{3, 5, 6, 3}
	expectedOutput := []int{0, 3}
	output := twoSum(input, 6)
	//fmt.Println(output)
	t.Log(output)
	if len(output) != len(expectedOutput) {
		t.Fail()
	}
	// if (output[0] != expectedOutput[0]) || (output[1] != expectedOutput[1]) {
	// 	t.Fail()
	// }
}
