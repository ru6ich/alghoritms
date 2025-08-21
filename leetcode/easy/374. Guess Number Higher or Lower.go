package easy

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	return 1
}

func guessNumber(n int) int {
	low := 0
	high := n

	for low <= high {
		mid := (low + high) / 2
		switch {
		case guess(mid) == 0:
			return mid
		case guess(mid) == -1:
			high = mid - 1
		default:
			low = mid + 1
		}

	}
	return -1
}
