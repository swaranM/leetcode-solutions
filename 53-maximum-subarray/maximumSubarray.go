package maximumsubarray

import "math"

func maxSubArray(nums []int) int {

	maxSum := nums[0]
	curSum := maxSum

	for i := 1; i < len(nums); i++ {
		curSum = int(math.Max(float64(nums[i]+curSum), float64(nums[i])))
		maxSum = int(math.Max(float64(curSum), float64(maxSum)))
	}

	return maxSum

}
