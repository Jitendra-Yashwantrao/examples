package mergesort

import (
	"examples/util"
	"fmt"
	"testing"
)

func TestMergeSort_1(t *testing.T) {

	list := []int{3, 7, 1}
	sortedList := MergeSort(list)
	fmt.Println("unsorted list", list)
	fmt.Println("Sorted List", sortedList)
	if len(list) != len(sortedList) {
		t.Error("length not matching")
	}
}

func TestMergeSort_10(t *testing.T) {

	list := util.GetRandomList(10)
	sortedList := MergeSort(list)
	fmt.Println("unsorted list", list)
	fmt.Println("Sorted List", sortedList)
	if len(list) != len(sortedList) {
		t.Error("length not matching")
	}
}

func TestMergeSort_19(t *testing.T) {

	list := util.GetRandomList(19)
	sortedList := MergeSort(list)
	fmt.Println("unsorted list", list)
	fmt.Println("Sorted List", sortedList)
	if len(list) != len(sortedList) {
		t.Error("length not matching")
	}
}
