package searchinrotatedsortedarray

func search(nums []int, target int) int {
	var left, right int
	var mid int

	right = len(nums) - 1

	for left <= right {
		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = left + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
