package mergesort

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2

	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func Merge(left, right []int) []int {
	length, i, j := len(left)+len(right), 0, 0

	result := make([]int, length, length)

	for k := 0; k < length; k++ {

		if i <= len(left)-1 && j > len(right)-1 {
			result[k] = left[i]
			i++

		} else if j <= len(right)-1 && i > len(left)-1 {
			result[k] = right[j]
			j++

		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {

			result[k] = right[j]
			j++
		}
	}
	return result
}
