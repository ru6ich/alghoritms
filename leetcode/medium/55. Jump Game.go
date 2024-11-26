package medium

func canJump(nums []int) bool {
	maxDistanceIdx := 0
	if len(nums) == 1 {
        return true
    }
	for i := 0; i < len(nums); i++ {
		if i > maxDistanceIdx {
            return false
        }
		if i + nums[i] > maxDistanceIdx {
			maxDistanceIdx = i + nums[i]
			if maxDistanceIdx >= len(nums) - 1 {
				return true
			}
		} 
	}
	return false
}
