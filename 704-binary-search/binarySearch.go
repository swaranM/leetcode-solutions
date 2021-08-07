package binarysearch

func search(nums []int, target int) int {

	l := 0
	h := len(nums) - 1
	var mid int

	for l <= h {
		mid = (l + h) / 2
		switch {
		case target == nums[mid]:
			return mid

		case target < nums[mid]:
			h = mid - 1

		case target > nums[mid]:
			l = mid + 1

		}

	}

	return -1
}
