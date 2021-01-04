package binarysearch

import "testing"

func TestBinarySearch_ItemPresent(t *testing.T) {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	output := BinarySearchItem(7, inputList)

	t.Log(output)

	if !output {

		t.Fail()

	}

}

func TestBinarySearch_ItemNotPresent(t *testing.T) {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	output := BinarySearchItem(17, inputList)

	t.Log(output)

	if output {

		t.Fail()

	}

}

func TestBinarySearch_EmptyList(t *testing.T) {
	inputList := []int{}

	output := BinarySearchItem(17, inputList)

	t.Log(output)

	if output {

		t.Fail()

	}

}
