package sortings

//find min element and insert to left
func SelectionSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
	return slice
}
