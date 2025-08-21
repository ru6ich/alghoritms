package easy

func findPeakElement(nums []int) int {
	low := 0
	high := len(nums) - 1

	for low < high {
		mid := (low + high) / 2
		switch {
		case nums[mid] < nums[mid+1]:
			low = mid + 1
		default:
			high = mid
		}
	}
	return low
}
