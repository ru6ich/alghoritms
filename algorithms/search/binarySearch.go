package binarySearch

func BinarySearch(arr [10]int, item int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if guess == item {
			return mid
		}

		if guess < item {
			low = mid + 1
		}

		if guess > item {
			high = mid - 1
		}
	}
	return -1
}
