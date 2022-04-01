package threesum

import (
	"fmt"
	"testing"
)

func TestThreeSum_1(t *testing.T) {

	//	list := util.GetRandomList(10)

	list := []int{-1, 0, 1, 2, -1, -4}

	tripletArray := ThreeSum(list)
	fmt.Println("list", list)
	fmt.Println("triplets", tripletArray)
	if len(tripletArray) != 2 {

		t.Error("Failed to find triplet array")
	}

	// 	Input: nums = [-1,0,1,2,-1,-4]
	// Output: [[-1,-1,2],[-1,0,1]]

}
