package sortings

func QuickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	pivot := slice[0]
	var less, greater []int
	for _, num := range slice[1:] {
		if num < pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(QuickSort(less), pivot)
	result = append(result, QuickSort(greater)...)
	return result
}
