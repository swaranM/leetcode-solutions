func maxProduct(nums []int) int {
	currMinProduct := nums[0]
	currMaxProduct := nums[0]
	maxProduct := nums[0]

	for _, v := range nums[1:] {

		if v == 0 {
			currMinProduct = 1
			currMaxProduct = 1
		}

		cmax := currMaxProduct * v
		cmin := currMinProduct * v

		currMaxProduct = int(math.Max(math.Max(float64(cmax), float64(cmin)), float64(v)))
		currMinProduct = int(math.Min(math.Min(float64(cmax), float64(cmin)), float64(v)))

		maxProduct = int(math.Max(float64(maxProduct), float64(currMaxProduct)))
	}

	return maxProduct
}