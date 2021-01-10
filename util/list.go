package util

import (
	"math/rand"
	"time"
)

//GetRandomList get random list of ints
func GetRandomList(size int) []int {
	rand.Seed(time.Now().UnixNano())
	newList := make([]int, size, size)
	for i := 0; i < size; i++ {
		newList[i] = rand.Intn(9) - rand.Intn(9)
	}
	return newList
}
