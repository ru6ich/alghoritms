package medium

func removeDuplicates(nums []int) int {
	k := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
