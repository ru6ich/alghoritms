package easy

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			high = mid - 1
		default:
			low = mid + 1
		}
	}
	return -1
}
